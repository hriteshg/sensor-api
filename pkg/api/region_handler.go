package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sensor-api/pkg/helpers"
	"sensor-api/pkg/model"
	"sensor-api/pkg/repository"
)

type RegionHandler struct {
	getMinimumTemperatureForRegion func(xMin float64, xMax float64, yMin float64, yMax float64, zMin float64, zMax float64) (float64, error)
	getMaximumTemperatureForRegion func(xMin float64, xMax float64, yMin float64, yMax float64, zMin float64, zMax float64) (float64, error)
}

// QueryMinTemperature is a handler that calculates minimum temperature in a given region.
// @Summary Calculate minimum temperature
// @Description Calculate minimum temperature inside a region
// @ID calculate-min-temperature
// @Success 200
// @Produce json
// @Param xMin query float64 true "Minimum x"
// @Param xMax query float64 true "Maximum x"
// @Param yMin query float64 true "Minimum y"
// @Param yMax query float64 true "Maximum y"
// @Param zMin query float64 true "Minimum z"
// @Param zMax query float64 true "Maximum z"
// @Router /region/temperature/min [get]
func (h RegionHandler) QueryMinTemperature(c *gin.Context) {
	xMin, err := helpers.GetFloatQueryParam(c, "xMin")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	xMax, err := helpers.GetFloatQueryParam(c, "xMax")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	yMin, err := helpers.GetFloatQueryParam(c, "yMin")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	yMax, err := helpers.GetFloatQueryParam(c, "yMax")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	zMin, err := helpers.GetFloatQueryParam(c, "zMin")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	zMax, err := helpers.GetFloatQueryParam(c, "zMax")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	r, err := h.getMinimumTemperatureForRegion(xMin, xMax, yMin, yMax, zMin, zMax)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, model.TemperatureResponse{
		Value: r,
		Scale: "Celsius",
	})
}

// QueryMaxTemperature is a handler that calculates maximum temperature in a given region.
// @Summary Calculate maximum temperature
// @Description Calculate maximum temperature inside a region
// @ID calculate-max-temperature
// @Success 200
// @Produce json
// @Param xMin query float64 true "Minimum x"
// @Param xMax query float64 true "Maximum x"
// @Param yMin query float64 true "Minimum y"
// @Param yMax query float64 true "Maximum y"
// @Param zMin query float64 true "Minimum z"
// @Param zMax query float64 true "Maximum z"
// @Router /region/temperature/max [get]
func (h RegionHandler) QueryMaxTemperature(c *gin.Context) {
	xMin, err := helpers.GetFloatQueryParam(c, "xMin")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	xMax, err := helpers.GetFloatQueryParam(c, "xMax")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	yMin, err := helpers.GetFloatQueryParam(c, "yMin")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	yMax, err := helpers.GetFloatQueryParam(c, "yMax")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	zMin, err := helpers.GetFloatQueryParam(c, "zMin")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	zMax, err := helpers.GetFloatQueryParam(c, "zMax")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	r, err := h.getMaximumTemperatureForRegion(xMin, xMax, yMin, yMax, zMin, zMax)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, model.TemperatureResponse{
		Value: r,
		Scale: "Celsius",
	})
}

func NewRegionHandler(db *gorm.DB) RegionHandler {
	sensorRepository := repository.NewSensorRepository(db)
	return RegionHandler{
		getMinimumTemperatureForRegion: sensorRepository.GetMinimumTemperatureForRegion,
		getMaximumTemperatureForRegion: sensorRepository.GetMaximumTemperatureForRegion,
	}
}
