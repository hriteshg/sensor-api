package repository

import (
	"database/sql"
	"gorm.io/gorm"
	"sensor-api/pkg/model"
	"sort"
	"time"
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

func (r SensorDataRepository) GetSensorAggregateForGroup(groupName string) (model.SensorGroupAggregate, error) {
	var result model.SensorGroupAggregate
	query := r.db.Table("sensor_groups").
		Joins("JOIN sensors ON sensor_groups.id = sensors.group_id").
		Joins("JOIN sensors_data ON sensors.id = sensors_data.sensor_id").
		Where("sensor_groups.name = ?", groupName).
		Select("AVG(sensors_data.transparency) as average_transparency, AVG(sensors_data.temperature) as average_temperature").
		Take(&result)
	result.Name = groupName
	return result, query.Error
}

func (r SensorDataRepository) GetSpeciesCountsForGroup(groupName string, topN *int, fromDateTime *time.Time, toDateTime *time.Time) ([]model.SpeciesCount, error) {
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
	if fromDateTime != nil && toDateTime != nil {
		for _, sp := range speciesLatestCounts {
			if sp.CreatedAt.After(*fromDateTime) && sp.CreatedAt.Before(*toDateTime) {
				filteredSpecies = append(filteredSpecies, sp)
			}
		}
	} else {
		filteredSpecies = speciesLatestCounts
	}

	sort.SliceStable(filteredSpecies, func(i, j int) bool {
		return filteredSpecies[i].Count > filteredSpecies[j].Count
	})

	var topSpeciesCounts []model.SpeciesCount

	if topN != nil && len(filteredSpecies) > *topN {
		topSpeciesCounts = filteredSpecies[:*topN]
	} else {
		topSpeciesCounts = filteredSpecies
	}

	return topSpeciesCounts, nil
}

func (r SensorDataRepository) GetMinimumTemperatureForRegion(xMin float64, xMax float64, yMin float64, yMax float64, zMin float64, zMax float64) (float64, error) {
	var minTemperature sql.NullFloat64

	query := r.db.Model(&model.SensorData{}).
		Joins("INNER JOIN sensors ON sensors.id = sensors_data.sensor_id").
		Where("sensors.x_coordinate >= ? AND sensors.x_coordinate <= ? AND sensors.y_coordinate >= ? AND sensors.y_coordinate <= ? AND sensors.z_coordinate >= ? AND sensors.z_coordinate <= ?",
			xMin, xMax, yMin, yMax, zMin, zMax).
		Select("MIN(sensors_data.temperature)").
		Scan(&minTemperature)

	if query.Error != nil {
		return 0, query.Error
	}
	if minTemperature.Valid {
		return minTemperature.Float64, query.Error
	} else {
		return 0, nil
	}
}

func (r SensorDataRepository) GetMaximumTemperatureForRegion(xMin float64, xMax float64, yMin float64, yMax float64, zMin float64, zMax float64) (float64, error) {
	var maxTemperature sql.NullFloat64

	query := r.db.Model(&model.SensorData{}).
		Joins("INNER JOIN sensors ON sensors.id = sensors_data.sensor_id").
		Where("sensors.x_coordinate >= ? AND sensors.x_coordinate <= ? AND sensors.y_coordinate >= ? AND sensors.y_coordinate <= ? AND sensors.z_coordinate >= ? AND sensors.z_coordinate <= ?",
			xMin, xMax, yMin, yMax, zMin, zMax).
		Select("MAX(sensors_data.temperature)").
		Scan(&maxTemperature)

	if query.Error != nil {
		return 0, query.Error
	}
	if maxTemperature.Valid {
		return maxTemperature.Float64, query.Error
	} else {
		return 0, nil
	}
}

func (r SensorDataRepository) GetAverageTemperatureForSensor(sensorId int64, fromDate time.Time, untilDateTime time.Time) (float64, error) {
	var averageTemperature sql.NullFloat64

	query := r.db.Model(&model.SensorData{}).
		Where("sensor_id = ? AND created_at >= ? AND created_at <= ?", sensorId, fromDate, untilDateTime).
		Select("AVG(temperature)").
		Scan(&averageTemperature)

	if query.Error != nil {
		return 0, query.Error
	}
	if averageTemperature.Valid {
		return averageTemperature.Float64, query.Error
	} else {
		return 0, nil
	}
}
