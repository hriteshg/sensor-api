package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"sensor-api/pkg/config"
	"sensor-api/pkg/db"
	"sensor-api/pkg/repository"
	"time"
)

func main() {
	c := config.Init()
	dbConfig := c.DBConfig()
	sensorsDB, err := db.OpenConn(dbConfig)
	if err != nil {
		log.Fatalf("Error connecting to DB %v", err)
	}
	r := repository.NewSensorGroupRepository(sensorsDB)
	ticker := time.Tick(10 * time.Second)

	for range ticker {
		groups, err := r.GetAllSensorGroups()
		if err != nil {
			return
		}
		for _, group := range groups {
			go makeRequest(fmt.Sprintf("http://localhost:3333/api//group/%s/transparency/average", group.Name))
			go makeRequest(fmt.Sprintf("http://localhost:3333/api//group/%s/temperature/average", group.Name))
		}

	}
}

func makeRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Received non-OK status code:", resp.StatusCode)
		return
	}

	// Process response body or perform other actions
	fmt.Println("GET request successful for URL:", url)
	log.Printf("GET request successful for URL: %v", resp.Body)
}
