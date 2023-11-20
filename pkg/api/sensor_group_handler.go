package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"sensor-api/pkg/cache"
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
	getFromCache               func(c *gin.Context, groupName string) (*model.SensorGroupAggregate, error)
	setInCache                 func(c *gin.Context, key string, value interface{}) error
}

func NewSensorGroupHandler(db *gorm.DB, cache cache.RedisCache) SensorGroupHandler {
	sensorRepository := repository.NewSensorRepository(db)

	return SensorGroupHandler{
		getGroup:                   sensorRepository.GetGroup,
		getSensorAggregateForGroup: sensorRepository.GetSensorAggregateForGroup,
		getSpeciesCountsForGroup:   sensorRepository.GetSpeciesCountsForGroup,
		getFromCache:               cache.GetFromCache,
		setInCache:                 cache.SetInCache,
	}
}

func (h SensorGroupHandler) QueryAverageTransparency(c *gin.Context) {
	groupName := c.Param("groupName")

	fromCache, err := h.getFromCache(c, groupName)
	if err == nil {
		log.Infof("Fetched from cache %v", fromCache)
		c.IndentedJSON(http.StatusOK, *fromCache)
		return
	}
	r, err := h.getSensorAggregateForGroup(groupName)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	marshal, err := json.Marshal(r)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = h.setInCache(c, groupName, marshal)
	if err != nil {
		log.Errorf("Set in cache returned error %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, r)
}

func (h SensorGroupHandler) QueryAverageTemperature(c *gin.Context) {
	groupName := c.Param("groupName")
	fromCache, err := h.getFromCache(c, groupName)
	if err == nil {
		log.Infof("Fetched from cache %v", fromCache)
		c.IndentedJSON(http.StatusOK, *fromCache)
		return
	}

	r, err := h.getSensorAggregateForGroup(groupName)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	marshal, err := json.Marshal(r)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = h.setInCache(c, groupName, marshal)
	if err != nil {
		log.Errorf("Set in cache returned error %v", err)
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
