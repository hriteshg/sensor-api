package main

import (
	"fmt"
	"log"
	"net/http"
	"sensor-api/pkg/api"
	"sensor-api/pkg/cache"
	"sensor-api/pkg/config"
	"sensor-api/pkg/db"
)

// @title Sensor API
// @version 1.0
// @description Sensor API to maintain sensor data and generate aggregate statistics.
// @host localhost:333
// @BasePath /api/v1
func main() {
	c := config.Init()
	dbConfig := c.DBConfig()
	sensorsDB, err := db.OpenConn(dbConfig)
	if err != nil {
		log.Fatalf("Error connecting to DB %v", err)
	}

	err = db.RunMigrations(sensorsDB, dbConfig, c.MigrationPath)
	if err != nil {
		log.Fatalf("Error running schema migration %v", err)
	}

	redisCache := cache.NewRedisCache(c.RedisUrl)
	server := api.NewRouter(sensorsDB, redisCache)
	log.Println(fmt.Sprintf("Listing for requests at http://localhost:%s/", c.PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", c.PORT), server))
}
