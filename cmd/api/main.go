package main

import (
	"fmt"
	"log"
	"net/http"
	"sensor-api/pkg/api"
	"sensor-api/pkg/config"
	"sensor-api/pkg/db"
)

func main() {
	c := config.Init()
	dbConfig := c.DBConfig()
	sensorsDB, err := db.OpenConn(dbConfig)
	if err != nil {
		log.Fatalf("Error connecting to DB %v", err)
	}
	if err != nil {
		log.Fatalf("Error connecting to DB %v", err)
	}
	err = db.RunMigrations(sensorsDB, dbConfig, c.MigrationPath)
	if err != nil {
		log.Fatalf("Error running schema migration %v", err)
	}

	server := api.NewRouter(sensorsDB)
	log.Println(fmt.Sprintf("Listing for requests at http://localhost:%s/", c.PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", c.PORT), server))
}