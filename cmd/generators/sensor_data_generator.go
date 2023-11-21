package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/rand"
	"sensor-api/pkg/config"
	"sensor-api/pkg/db"
	"sensor-api/pkg/model"
	"sensor-api/pkg/repository"
	"time"
)

func main() {
	c := config.Init()
	dbConfig := c.DBConfig()
	sensorsDB, err := db.OpenConn(dbConfig)
	if err != nil {
		log.Fatalf("Error connecting to DB %v", err)
	}

	generator := NewSensorDataGenerator(sensorsDB)
	generator.Start()
}

type SensorDataGenerator struct {
	getAllSensors    func() ([]model.Sensor, error)
	createSensorData func(data model.SensorData) error
}

func NewSensorDataGenerator(db *gorm.DB) SensorDataGenerator {
	r := repository.NewSensorRepository(db)
	s := repository.NewSensorDataRepository(db)
	return SensorDataGenerator{
		getAllSensors:    r.GetAllSensors,
		createSensorData: s.CreateSensorData,
	}
}

func (s SensorDataGenerator) Start() {
	sensors, err := s.getAllSensors()
	if err != nil {
		return
	}
	for _, sensor := range sensors {
		go s.runScheduler(sensor)
	}

	select {}
}

var predefinedFishes = []string{"Species A", "Species B", "Species C"}
var sensorDataCh = make(chan model.SensorData)

func (s SensorDataGenerator) runScheduler(sensor model.Sensor) {
	source := rand.NewSource(sensor.ID) // Create a new source for random number generation

	for {
		select {
		case <-time.After(time.Duration(sensor.DataOutputRate)):

			data := generateSensorData(sensor, source)
			// Send data to processing or storage
			sensorDataCh <- data
		}
	}
}

func (s SensorDataGenerator) processSensorData(sensorDataCh <-chan model.SensorData) {
	for data := range sensorDataCh {
		err := s.createSensorData(data)
		if err != nil {
			return
		}
		fmt.Printf("Writing sensor data to the database: %+v\n", data)
	}
}

func generateTemperature(depthCoefficient float64, source rand.Source) float64 {
	r := rand.New(source)
	return r.Float64() * depthCoefficient * 100 // Adjust the range (0-100) as needed
}

func generateTransparency(previousTransparency int64, source rand.Source) int64 {
	maxDiff := int64(10)
	r := rand.New(source)

	if previousTransparency < 1 {
		previousTransparency = 1
	} else if previousTransparency > 100 {
		previousTransparency = 100
	}

	return previousTransparency + r.Int63n(maxDiff*2) - maxDiff
}

func generateSensorData(sensor model.Sensor, source rand.Source) model.SensorData {
	temperature := generateTemperature(sensor.ZCoordinate, source)
	transparency := generateTransparency(0, source)

	var fishData []model.FishData
	for _, fishName := range selectRandomFishSpecies(2) {
		fish := model.FishData{
			SpeciesName:  fishName,
			Count:        rand.Int63n(10),
			SensorDataID: sensor.ID,
		}
		fishData = append(fishData, fish)
	}

	return model.SensorData{
		Transparency: transparency,
		Temperature:  temperature,
		SensorID:     sensor.ID,
		FishData:     fishData,
		CreatedAt:    time.Now(),
	}
}

func selectRandomFishSpecies(count int) []string {
	var selectedFish []string
	availableFishes := len(predefinedFishes)

	if count > availableFishes {
		count = availableFishes
	}

	rand.Shuffle(len(predefinedFishes), func(i, j int) {
		predefinedFishes[i], predefinedFishes[j] = predefinedFishes[j], predefinedFishes[i]
	})

	for i := 0; i < count; i++ {
		selectedFish = append(selectedFish, predefinedFishes[i])
	}

	return selectedFish
}
