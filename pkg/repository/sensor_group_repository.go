package repository

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"sensor-api/pkg/model"
)

type SensorGroupRepository struct {
	db *gorm.DB
}

func NewSensorGroupRepository(db *gorm.DB) SensorGroupRepository {
	return SensorGroupRepository{
		db: db,
	}
}

func (r SensorGroupRepository) CreateSensors(groupNames []string) {
	for _, groupName := range groupNames {
		var sensors []model.Sensor
		for i := 0; i < 3; i++ {
			xVal := rand.Float64() * 100
			yVal := rand.Float64() * 100
			zVal := rand.Float64() * 100
			sensors = append(sensors, model.Sensor{
				Name:           fmt.Sprintf("Sensor-%s-%d", groupName, i+1),
				XCoordinate:    xVal,
				YCoordinate:    yVal,
				ZCoordinate:    zVal,
				DataOutputRate: rand.Int63n(100),
			})
		}
		group := model.SensorGroup{Name: groupName, Sensors: sensors}
		r.db.Save(&group)
	}
}

func (r SensorGroupRepository) GetAllSensorGroups() ([]model.SensorGroup, error) {
	var results []model.SensorGroup
	tx := r.db.Find(&results)
	return results, tx.Error
}
