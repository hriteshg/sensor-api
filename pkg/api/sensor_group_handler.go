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

// QueryAverageTransparency is a handler that retrieves average transparency of sensors within a sensor group.
// @Summary Collect average transparency of sensors within a sensor group
// @Description Collect average transparency
// @ID query-average-transparency
// @Param groupName path string true "Group name"
// @Produce json
// @Success 200
// @Router /group/:groupName/transparency [get]
func (h SensorGroupHandler) QueryAverageTransparency(c *gin.Context) {
	groupName := c.Param("groupName")
	if len(groupName) == 0 {
		c.IndentedJSON(http.StatusBadRequest, "Invalid group name")
		return
	}

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

// QueryAverageTemperature is a handler that retrieves average temperature of sensors within a sensor group.
// @Summary Collect average temperature of sensors within a sensor group
// @Description Collect average temperature
// @ID query-average-temperature
// @Param groupName path string true "Group name"
// @Produce json
// @Success 200
// @Router /group/:groupName/temperature [get]
func (h SensorGroupHandler) QueryAverageTemperature(c *gin.Context) {
	groupName := c.Param("groupName")
	if len(groupName) == 0 {
		c.IndentedJSON(http.StatusBadRequest, "Invalid group name")
		return
	}

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

// QuerySpecies is a handler that retrieves full list of species with counts currently detected inside the group.
// @Summary Retrieve full list of species inside the group
// @Description Retrieves full list of species with counts currently detected inside the group.
// @ID query-species
// @Param groupName path string true "Group name"
// @Produce json
// @Success 200
// @Router /group/:groupName/species [get]
func (h SensorGroupHandler) QuerySpecies(c *gin.Context) {
	groupName := c.Param("groupName")
	if len(groupName) == 0 {
		c.IndentedJSON(http.StatusBadRequest, "Invalid group name")
		return
	}

	r, err := h.getSpeciesCountsForGroup(groupName, nil, nil, nil)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, r)
}

// QueryTopNSpecies is a handler that retrieves list of top N species with counts currently detected inside the group.
// @Summary Retrieve list of top N species inside the group
// @Description Retrieves list of top N species with counts currently detected inside the group.
// @ID query-species-with-filter
// @Param groupName path string true "Group name"
// @Param N path int true "Top N species count"
// @Param from query int64 false "From time in Unix timestamp"
// @Param till query int64 false "Till time in Unix timestamp"
// @Produce json
// @Success 200
// @Router /group/:groupName/species/top/:N [get]
func (h SensorGroupHandler) QueryTopNSpecies(c *gin.Context) {
	groupName := c.Param("groupName")
	if len(groupName) == 0 {
		c.IndentedJSON(http.StatusBadRequest, "Invalid group name")
		return
	}

	topN, err := strconv.Atoi(c.Param("N"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Invalid N")
		return
	}

	from, err := helpers.GetTime(c, "fromDateTime", nil)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Invalid fromDateTime")
		return
	}

	to, err := helpers.GetTime(c, "untilDateTime", nil)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Invalid untilDateTime")
		return
	}

	r, err := h.getSpeciesCountsForGroup(groupName, &topN, from, to)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, r)
}
