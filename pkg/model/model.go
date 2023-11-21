package model

import (
	"errors"
	"github.com/gin-gonic/gin"
	"sensor-api/pkg/helpers"
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
	Name                string  `json:"name"`
	AverageTemperature  float64 `json:"average_temperature"`
	AverageTransparency float64 `json:"average_transparency"`
}

type SpeciesCount struct {
	Name      string    `json:"name"`
	Count     int64     `json:"count"`
	CreatedAt time.Time `json:"created_at"`
}

type TemperatureResponse struct {
	Value float64 `json:"value"`
	Scale string  `json:"scale"`
}

type Region struct {
	XMin float64
	XMax float64
	YMin float64
	YMax float64
	ZMin float64
	ZMax float64
}

func BuildRegion(c *gin.Context) (*Region, error) {
	xMin, err := helpers.GetFloatQueryParam(c, "xMin")
	if err != nil {
		return nil, errors.New("invalid xMin")
	}
	xMax, err := helpers.GetFloatQueryParam(c, "xMax")
	if err != nil {
		return nil, errors.New("invalid xMax")
	}
	yMin, err := helpers.GetFloatQueryParam(c, "yMin")
	if err != nil {
		return nil, errors.New("invalid yMin")
	}
	yMax, err := helpers.GetFloatQueryParam(c, "yMax")
	if err != nil {
		return nil, errors.New("invalid yMax")
	}
	zMin, err := helpers.GetFloatQueryParam(c, "zMin")
	if err != nil {
		return nil, errors.New("invalid zMin")
	}
	zMax, err := helpers.GetFloatQueryParam(c, "zMax")
	if err != nil {
		return nil, errors.New("invalid zMax")
	}
	return &Region{
		XMin: xMin,
		XMax: xMax,
		YMin: yMin,
		YMax: yMax,
		ZMin: zMin,
		ZMax: zMax,
	}, nil
}
