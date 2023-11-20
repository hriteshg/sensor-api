package main

import (
	log "github.com/sirupsen/logrus"
	"sensor-api/pkg/config"
	"sensor-api/pkg/db"
	"sensor-api/pkg/repository"
)

func main() {
	c := config.Init()
	dbConfig := c.DBConfig()
	sensorsDB, err := db.OpenConn(dbConfig)
	if err != nil {
		log.Fatalf("Error connecting to DB %v", err)
	}
	groupNames := []string{
		"alpha",
		"beta",
		"gamma",
	}

	groupRepository := repository.NewSensorGroupRepository(sensorsDB)
	groupRepository.CreateSensors(groupNames)
}
