package repository

import (
	"gorm.io/gorm"
	"sensor-api/pkg/model"
)

type SensorDataRepository struct {
	db *gorm.DB
}

func NewSensorDataRepository(db *gorm.DB) SensorDataRepository {
	return SensorDataRepository{
		db: db,
	}
}

func (r SensorDataRepository) CreateSensorData(data model.SensorData) error {
	tx := r.db.Save(&data)
	return tx.Error
}
