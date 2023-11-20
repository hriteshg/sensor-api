package model

import (
	"time"
)

type SensorGroup struct {
	ID      int64    `gorm:"primaryKey"`
	Name    string   `gorm:"unique"`
	Sensors []Sensor `gorm:"foreignKey:GroupID"`
}

func (SensorGroup) TableName() string {
	return "sensor_groups"
}

type Sensor struct {
	ID             int64  `gorm:"primaryKey"`
	Name           string `gorm:"unique"`
	GroupID        int64
	XCoordinate    float64
	YCoordinate    float64
	ZCoordinate    float64
	DataOutputRate int64
}

func (Sensor) TableName() string {
	return "sensors"
}

type SensorData struct {
	ID           int64 `gorm:"primaryKey"`
	Transparency int64
	Temperature  float64
	SensorID     int64
	FishData     []FishData `gorm:"foreignKey:SensorDataID"`
	CreatedAt    time.Time  `gorm:"type:TIMESTAMP WITH TIME ZONE;default:CURRENT_TIMESTAMP"`
}

func (SensorData) TableName() string {
	return "sensors_data"
}

type FishData struct {
	ID           int64 `gorm:"primaryKey"`
	SpeciesName  string
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

type SpeciesCount struct {
	Name      string
	Count     int64
	CreatedAt time.Time
}
