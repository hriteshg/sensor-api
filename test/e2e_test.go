package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sensor-api/pkg/model"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const BaseUrl = "http://localhost:3333/api/v1"

func TestAverageTransparencyApi(t *testing.T) {
	var averageTransparency model.SensorGroupAggregate
	response, err := http.Get(BaseUrl + "/group/alpha/transparency/average")
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &averageTransparency)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "alpha", averageTransparency.Name)
	assert.Equal(t, float64(15), averageTransparency.AverageTransparency)
}

func TestAverageTemperatureApi(t *testing.T) {
	var averageTemperature model.SensorGroupAggregate
	response, err := http.Get(BaseUrl + "/group/alpha/temperature/average")
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &averageTemperature)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, float64(45), averageTemperature.AverageTemperature)
}

func TestTotalSpeciesInGroupApi(t *testing.T) {
	var species []model.SpeciesCount
	response, err := http.Get(BaseUrl + "/group/alpha/species")
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &species)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 2, len(species))
	assert.Equal(t, "Pacific Cod", species[0].Name)
	assert.Equal(t, int64(12), species[0].Count)
}

func TestTopNSpeciesInGroupApi(t *testing.T) {
	var species []model.SpeciesCount
	response, err := http.Get(BaseUrl + "/group/alpha/species/top/1")
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &species)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 1, len(species))
	assert.Equal(t, int64(12), species[0].Count)
}

func TestTopNSpeciesInTheGroupBetweenApi(t *testing.T) {
	var species []model.SpeciesCount
	response, err := http.Get(BaseUrl + "/group/Group2/species/top/2?from=1700137279&till=1700137279")
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &species)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, len(species), 0)
}

func TestMinTemperatureInGroupApi(t *testing.T) {
	var minTemperature model.TemperatureResponse
	response, err := http.Get(BaseUrl + "/region/temperature/min?xMin=10&xMax=48&yMin=2&yMax=100&zMin=60&zMax=100")
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &minTemperature)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, float64(50), minTemperature.Value)
}

func TestMaxTemperatureInGroupApi(t *testing.T) {
	var maxTemperature model.TemperatureResponse
	response, err := http.Get(BaseUrl + "/region/temperature/max?xMin=10&xMax=48&yMin=2&yMax=100&zMin=60&zMax=100")
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &maxTemperature)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, float64(70), maxTemperature.Value)
}

func TestAverageTemperatureBySensorApi(t *testing.T) {
	var averageTemperature model.TemperatureResponse
	currentTimestamp := int64(time.Now().Unix())
	response, err := http.Get(BaseUrl + "/sensor/alpha 1/temperature/average?from=1697438904&till=" + strconv.FormatInt(currentTimestamp, 10))
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &averageTemperature)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, float64(40), averageTemperature.Value)
}
