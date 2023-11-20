package model

import (
	"math/big"
	"time"
)

type SensorGroup struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

func (SensorGroup) TableName() string {
	return "sensor_groups"
}

type Sensor struct {
	ID             int64  `gorm:"primaryKey"`
	Name           string `gorm:"unique"`
	GroupID        int64
	x              big.Float
	y              big.Float
	z              big.Float
	DataOutputRate int64
}

func (Sensor) TableName() string {
	return "sensors"
}

type SensorData struct {
	ID           int64 `gorm:"primaryKey"`
	Transparency int64
	Temperature  int64
	SensorID     uint
	FishData     []FishData `gorm:"foreignKey:SensorDataID"`
	CreatedAt    time.Time  `gorm:"type:TIMESTAMP WITH TIME ZONE;default:CURRENT_TIMESTAMP"`
}

type SpeciesCount struct {
	Name      string
	Count     int64
	CreatedAt time.Time
}

func (SensorData) TableName() string {
	return "sensors_data"
}

type FishSpecies struct {
	ID          int64 `gorm:"primaryKey"`
	SpeciesName string
}

func (FishSpecies) TableName() string {
	return "fish_species"
}

type FishData struct {
	ID           int64 `gorm:"primaryKey"`
	SpeciesID    int64
	SensorDataID int64
	Count        int64
}

func (FishData) TableName() string {
	return "fish_data"
}

type SensorGroupAggregate struct {
	Name                string
	AverageTemperature  float64
	AverageTransparency float64
}
