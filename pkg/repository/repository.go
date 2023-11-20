package repository

import (
	"errors"
	"gorm.io/gorm"
	"sensor-api/pkg/model"
	"sort"
	"time"
)

type SensorRepository struct {
	db *gorm.DB
}

func NewSensorRepository(db *gorm.DB) SensorRepository {
	return SensorRepository{
		db: db,
	}
}

func (r SensorRepository) GetGroup() (model.SensorData, error) {
	var c model.SensorData

	result := r.db.Preload("FishData").Find(&c, "id = ?", 1)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.SensorData{}, nil
	}

	return c, result.Error
}

func (r SensorRepository) GetSensorAggregateForAGroup(groupName string) (model.SensorGroupAggregate, error) {
	var result model.SensorGroupAggregate
	query := r.db.Table("sensor_groups").
		Joins("JOIN sensors ON sensor_groups.id = sensors.group_id").
		Joins("JOIN sensors_data ON sensors.id = sensors_data.sensor_id").
		Where("sensor_groups.name = ?", groupName).
		Select("AVG(sensors_data.transparency) as avg_transparency, AVG(sensors_data.temperature) as avg_temperature").
		Take(&result)
	result.Name = groupName
	return result, query.Error
}

func (r SensorRepository) GetSpeciesCountsForGroup(groupName string, topN *int, fromDateTime *time.Time, toDateTime *time.Time) ([]model.SpeciesCount, error) {
	var speciesLatestCounts []model.SpeciesCount

	r.db.Table("sensor_groups").
		Joins("JOIN sensors ON sensor_groups.id = sensors.group_id").
		Joins("JOIN sensors_data ON sensors.id = sensors_data.sensor_id").
		Joins("JOIN fish_data ON sensors_data.id = fish_data.sensor_data_id").
		Where("sensor_groups.name = ?", groupName).
		Select("fish_data.species_name as name, fish_data.count as count, fish_data.created_at as created_at").
		Where("(fish_data.species_name, fish_data.created_at) IN (SELECT fd.species_name, MAX(fd.created_at) FROM fish_data fd GROUP BY fd.species_name)").
		Scan(&speciesLatestCounts)

	var filteredSpecies []model.SpeciesCount
	for _, sp := range speciesLatestCounts {
		if sp.CreatedAt.After(*fromDateTime) && sp.CreatedAt.Before(*toDateTime) {
			filteredSpecies = append(filteredSpecies, sp)
		}
	}
	sort.SliceStable(filteredSpecies, func(i, j int) bool {
		return filteredSpecies[i].Count > filteredSpecies[j].Count
	})

	var topSpeciesCounts []model.SpeciesCount
	if len(filteredSpecies) > *topN {
		topSpeciesCounts = filteredSpecies[:*topN]
	} else {
		topSpeciesCounts = filteredSpecies
	}
	return topSpeciesCounts, nil
}

func (r SensorRepository) GetMinimumTemperatureForRegion(xMin float64, xMax float64, yMin float64, yMax float64, zMin float64, zMax float64) (float64, error) {
	var minTemperature float64

	query := r.db.Model(&model.SensorData{}).
		Joins("INNER JOIN sensors ON sensors.id = sensors_data.sensor_id").
		Where("sensors.x >= ? AND sensors.x <= ? AND sensors.y >= ? AND sensors.y <= ? AND sensors.z >= ? AND sensors.z <= ?",
			xMin, xMax, yMin, yMax, zMin, zMax).
		Select("MIN(sensors_data.temperature)").
		Scan(&minTemperature)

	return minTemperature, query.Error
}

func (r SensorRepository) GetMaximumTemperatureForRegion(xMin float64, xMax float64, yMin float64, yMax float64, zMin float64, zMax float64) (float64, error) {
	var minTemperature float64

	query := r.db.Model(&model.SensorData{}).
		Joins("INNER JOIN sensors ON sensors.id = sensors_data.sensor_id").
		Where("sensors.x >= ? AND sensors.x <= ? AND sensors.y >= ? AND sensors.y <= ? AND sensors.z >= ? AND sensors.z <= ?",
			xMin, xMax, yMin, yMax, zMin, zMax).
		Select("MAX(sensors_data.temperature)").
		Scan(&minTemperature)

	return minTemperature, query.Error
}

func (r SensorRepository) GetAverageTemperatureForSensor(sensorId int64, fromDate time.Time, untilDateTime time.Time) (float64, error) {
	var averageTemperature float64

	query := r.db.Model(&model.SensorData{}).
		Where("sensor_id = ? AND created_at >= ? AND created_at <= ?", sensorId, fromDate, untilDateTime).
		Select("AVG(temperature)").
		Scan(&averageTemperature)

	return averageTemperature, query.Error
}
