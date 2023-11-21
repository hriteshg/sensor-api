package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sensor-api/pkg/helpers"
	"sensor-api/pkg/model"
	"sensor-api/pkg/repository"
	"strings"
	"time"
)

type SensorHandler struct {
	getAverageTemperatureForSensor func(sensorId int64, fromDate time.Time, untilDateTime time.Time) (float64, error)
	getSensorByCodeName            func(codeName string) (model.Sensor, error)
}

func NewSensorHandler(db *gorm.DB) SensorHandler {
	sensorRepository := repository.NewSensorDataRepository(db)
	s := repository.NewSensorRepository(db)
	return SensorHandler{
		getSensorByCodeName:            s.GetSensorByCodeName,
		getAverageTemperatureForSensor: sensorRepository.GetAverageTemperatureForSensor,
	}
}

// QueryAverageTemperature is a handler that calculates average temperature in a given time interval by a sensor.
// @Summary Calculates average temperature
// @Description Calculate average temperature in a given time interval by a sensor
// @ID calculate-average-temperature-by-sensor
// @Produce json
// @Success 200
// @Param from query int64 true "From time in Unix timestamp"
// @Param till query int64 true "Till time in Unix timestamp"
// @Param codeName path string  true "Code name of the sensor"
// @Router /sensor/:codeName/temperature/average [get]
func (h SensorHandler) QueryAverageTemperature(c *gin.Context) {
	codeName, err := h.getCodeName(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	from, err := helpers.GetTime(c, "from", nil)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	to, err := helpers.GetTime(c, "till", nil)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	sensor, err := h.getSensorByCodeName(codeName)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	r, err := h.getAverageTemperatureForSensor(sensor.ID, *from, *to)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, model.TemperatureResponse{
		Value: r,
		Scale: "Celsius",
	})
}

func (h SensorHandler) getCodeName(c *gin.Context) (string, error) {
	codeName := c.Param("codeName")
	if len(codeName) == 0 {
		return "", errors.New("invalid code name")
	}
	parts := strings.Split(codeName, " ")
	if len(parts) != 2 {
		return "", errors.New("invalid code name")
	}
	return codeName, nil
}
