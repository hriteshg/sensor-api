package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sensor-api/pkg/helpers"
	"sensor-api/pkg/repository"
	"strconv"
	"time"
)

type SensorHandler struct {
	getAverageTemperatureForSensor func(sensorId int64, fromDate time.Time, untilDateTime time.Time) (float64, error)
}

func (h SensorHandler) QueryAverageTemperature(c *gin.Context) {
	codeName := c.Param("codeName")
	from, err := helpers.GetDate(c, "fromDate", nil)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	to, err := helpers.GetDate(c, "untilDateTime", nil)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	atoi, _ := strconv.Atoi(codeName)

	r, err := h.getAverageTemperatureForSensor(int64(atoi), *from, *to)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, r)
}

func NewSensorHandler(db *gorm.DB) SensorHandler {
	sensorRepository := repository.NewSensorRepository(db)
	return SensorHandler{
		getAverageTemperatureForSensor: sensorRepository.GetAverageTemperatureForSensor,
	}
}
