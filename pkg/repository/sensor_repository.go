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

func (r SensorRepositoryOne) GetSensorByCodeName(codeName string) (model.Sensor, error) {
	var sensor model.Sensor
	query := r.db.Where("name = ?", codeName).First(&sensor)
	return sensor, query.Error

}
