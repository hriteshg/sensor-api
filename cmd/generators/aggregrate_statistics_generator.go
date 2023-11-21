package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
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
	r := repository.NewSensorGroupRepository(sensorsDB)
	ticker := time.Tick(10 * time.Second)

	for range ticker {
		groups, err := r.GetAllSensorGroups()
		if err != nil {
			return
		}
		for _, group := range groups {
			go getAverageTransparency(group)
			go getAverageTemperature(group)
		}
	}
}

func getAverageTransparency(sensorGroup model.SensorGroup) {
	a := makeRequest(fmt.Sprintf("http://localhost:3333/api/v1/group/%s/transparency/average", sensorGroup.Name))
	if a == nil {
		return
	}
	log.Printf("Average Transparency for Sensor Group: %v is %v", sensorGroup.Name, a.AverageTransparency)
	log.Printf("--------------")
}

func getAverageTemperature(sensorGroup model.SensorGroup) {
	a := makeRequest(fmt.Sprintf("http://localhost:3333/api/v1/group/%s/temperature/average", sensorGroup.Name))
	if a == nil {
		return
	}
	log.Printf("Average Temperature for Sensor Group: %v is %v", sensorGroup.Name, a.AverageTemperature)
	log.Printf("--------------")
}

func makeRequest(url string) *model.SensorGroupAggregate {
	var aggregateModel model.SensorGroupAggregate
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error making GET request:", err)
		return nil
	}
	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &aggregateModel)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return nil
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Println("Received non-OK status code:", resp.StatusCode)
		return nil
	}

	return &aggregateModel
}
