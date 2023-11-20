package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"net/http"

	_ "sensor-api/docs"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rtr := gin.New()
	rtr.Use(gin.Recovery())
	rtr.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	rtr.GET("/health", checkHealth)
	addRoutes(rtr, db)
	return rtr
}

func addRoutes(rtr *gin.Engine, db *gorm.DB) {
	sensorGroupHandler := NewSensorGroupHandler(db)
	sensorHandler := NewSensorHandler(db)
	regionHandler := NewRegionHandler(db)

	rtr.GET("/api/group/:groupName/transparency/average", sensorGroupHandler.QueryAverageTransparency)
	rtr.GET("/api/group/:groupName/temperature/average", sensorGroupHandler.QueryAverageTemperature)
	rtr.GET("/api/group/:groupName/species", sensorGroupHandler.QuerySpecies)
	rtr.GET("/api/group/:groupName/species/top/:N", sensorGroupHandler.QueryTopNSpecies)
	rtr.GET("/api/region/temperature/min", regionHandler.QueryMinTemperature)
	rtr.GET("/api/region/temperature/max", regionHandler.QueryMaxTemperature)
	rtr.GET("/api/sensor/:codeName/temperature/average", sensorHandler.QueryAverageTemperature)
}

func checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
