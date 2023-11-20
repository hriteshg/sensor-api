package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sensor-api/pkg/helpers"
	"sensor-api/pkg/repository"
)

type RegionHandler struct {
	getMinimumTemperatureForRegion func(xMin float64, xMax float64, yMin float64, yMax float64, zMin float64, zMax float64) (float64, error)
	getMaximumTemperatureForRegion func(xMin float64, xMax float64, yMin float64, yMax float64, zMin float64, zMax float64) (float64, error)
}

func (h RegionHandler) QueryMinTemperature(c *gin.Context) {
	xMin, err := helpers.GetFloatQueryParam(c, "xMin")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	xMax, err := helpers.GetFloatQueryParam(c, "xMax")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	yMin, err := helpers.GetFloatQueryParam(c, "yMin")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	yMax, err := helpers.GetFloatQueryParam(c, "yMax")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	zMin, err := helpers.GetFloatQueryParam(c, "zMin")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	zMax, err := helpers.GetFloatQueryParam(c, "zMax")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	r, err := h.getMinimumTemperatureForRegion(xMin, xMax, yMin, yMax, zMin, zMax)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, r)
}

func (h RegionHandler) QueryMaxTemperature(c *gin.Context) {
	xMin, err := helpers.GetFloatQueryParam(c, "xMin")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	xMax, err := helpers.GetFloatQueryParam(c, "xMax")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	yMin, err := helpers.GetFloatQueryParam(c, "yMin")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	yMax, err := helpers.GetFloatQueryParam(c, "yMax")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	zMin, err := helpers.GetFloatQueryParam(c, "zMin")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	zMax, err := helpers.GetFloatQueryParam(c, "zMax")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	r, err := h.getMaximumTemperatureForRegion(xMin, xMax, yMin, yMax, zMin, zMax)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, r)
}

func NewRegionHandler(db *gorm.DB) RegionHandler {
	sensorRepository := repository.NewSensorRepository(db)
	return RegionHandler{
		getMinimumTemperatureForRegion: sensorRepository.GetMinimumTemperatureForRegion,
		getMaximumTemperatureForRegion: sensorRepository.GetMaximumTemperatureForRegion,
	}
}
