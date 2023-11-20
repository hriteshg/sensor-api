package repository

import (
	"gorm.io/gorm"
	"sensor-api/pkg/model"
)

type SensorRepositoryOne struct {
	db *gorm.DB
}

func NewSensorRepositoryOne(db *gorm.DB) SensorRepositoryOne {
	return SensorRepositoryOne{
		db: db,
	}
}

func (r SensorRepositoryOne) GetAllSensors() ([]model.Sensor, error) {
	var results []model.Sensor
	tx := r.db.Find(&results)
	return results, tx.Error
}
