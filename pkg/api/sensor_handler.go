package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sensor-api/pkg/helpers"
	"sensor-api/pkg/model"
	"sensor-api/pkg/repository"
	"time"
)

type SensorHandler struct {
	getAverageTemperatureForSensor func(sensorId int64, fromDate time.Time, untilDateTime time.Time) (float64, error)
	getSensorByCodeName            func(codeName string) (model.Sensor, error)
}

func (h SensorHandler) QueryAverageTemperature(c *gin.Context) {
	codeName := c.Param("codeName")
	from, err := helpers.GetTime(c, "from", nil)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	to, err := helpers.GetTime(c, "till", nil)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	sensor, err := h.getSensorByCodeName(codeName)
	if err != nil {
		return
	}

	r, err := h.getAverageTemperatureForSensor(sensor.ID, *from, *to)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, model.TemperatureResponse{
		Value: r,
		Scale: "Celsius",
	})
}

func NewSensorHandler(db *gorm.DB) SensorHandler {
	sensorRepository := repository.NewSensorRepository(db)
	s := repository.NewSensorRepositoryOne(db)
	return SensorHandler{
		getSensorByCodeName:            s.GetSensorByCodeName,
		getAverageTemperatureForSensor: sensorRepository.GetAverageTemperatureForSensor,
	}
}
