package repository

import (
	"gorm.io/gorm"
	"sensor-api/pkg/model"
)

type SensorRepository struct {
	db *gorm.DB
}

func NewSensorRepository(db *gorm.DB) SensorRepository {
	return SensorRepository{
		db: db,
	}
}

func (r SensorRepository) GetAllSensors() ([]model.Sensor, error) {
	var results []model.Sensor
	tx := r.db.Find(&results)
	return results, tx.Error
}

func (r SensorRepository) GetSensorByCodeName(codeName string) (model.Sensor, error) {
	var sensor model.Sensor
	query := r.db.Where("name = ?", codeName).First(&sensor)
	return sensor, query.Error

}
