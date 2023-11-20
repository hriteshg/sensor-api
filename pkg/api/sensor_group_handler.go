package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sensor-api/pkg/helpers"
	"sensor-api/pkg/model"
	"sensor-api/pkg/repository"
	"strconv"
	"time"
)

type SensorGroupHandler struct {
	getGroup                   func() (model.SensorData, error)
	getSensorAggregateForGroup func(groupName string) (model.SensorGroupAggregate, error)
	getSpeciesCountsForGroup   func(groupName string, topN *int, fromDateTime *time.Time, toDateTime *time.Time) ([]model.SpeciesCount, error)
}

func NewSensorGroupHandler(db *gorm.DB) SensorGroupHandler {
	sensorRepository := repository.NewSensorRepository(db)
	return SensorGroupHandler{
		getGroup:                   sensorRepository.GetGroup,
		getSensorAggregateForGroup: sensorRepository.GetSensorAggregateForGroup,
		getSpeciesCountsForGroup:   sensorRepository.GetSpeciesCountsForGroup,
	}
}

func (h SensorGroupHandler) QueryAverageTransparency(c *gin.Context) {
	groupName := c.Param("groupName")
	r, err := h.getSensorAggregateForGroup(groupName)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, r)
}

func (h SensorGroupHandler) QueryAverageTemperature(c *gin.Context) {
	groupName := c.Param("groupName")
	r, err := h.getSensorAggregateForGroup(groupName)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, r)
}

func (h SensorGroupHandler) QuerySpecies(c *gin.Context) {
	groupName := c.Param("groupName")
	r, err := h.getSpeciesCountsForGroup(groupName, nil, nil, nil)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, r)
}

func (h SensorGroupHandler) QueryTopNSpecies(c *gin.Context) {
	groupName := c.Param("groupName")
	topN, _ := strconv.Atoi(c.Param("N"))
	from, err := helpers.GetTime(c, "fromDateTime", nil)
	if err != nil {
		return
	}
	to, err := helpers.GetTime(c, "untilDateTime", nil)
	if err != nil {
		return
	}
	r, err := h.getSpeciesCountsForGroup(groupName, &topN, from, to)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, r)
}
