## Sensor API

This repository is backend implementation of Sensor API which collects Sensor data and calculates various aggregation
statistics.
The API is built using the Go language and the Gin framework.

## Features

* Captures Sensor group and Sensors information
* Collects data from sensors
* Performs aggregations on sensor data

## Tech Stack

* GoLang
* PostgresSQL for data storage
* Docker
* Swagger API documentation
* Redis

## Prerequisites

Ensure the below are installed:

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)

## Local Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/hriteshg/sensor-api.git
    cd sensor-api
    ```

2. Rename `env.example` to  `.env` file in the project root and configure the data

## Run the Service

Use docker and run the current Go application using below command:

    ```bash
         docker-compose up --build
    ```

Alternatively, you can also run the Go run separately using below commands:

    ```bash
         docker-compose up
         go run ./cmd/api/main.go
    ```

The API will be available at http://localhost:3333

## Run the Generators

There are 3 generators available.

* **Sensors Generator** - To seed sensor group and sensors data (**sensor_data_generator.go**)
* **Sensor Data Generator** - To seed sensor data based on sensor's data output rate (**sensor_data_generator.go**)
* **Aggregate Statistics Generator** = To fetch required statistics (**aggregate_statistics_generator.go**)

To run the required generator, use below command:

    ```bash
         docker-compose up
         go run ./cmd/generators/<generatorName>.go
    ```

## Swagger Documentation

To generate Swagger docs, use below command:

    ```bash
         swag init -g ../../cmd/api/main.go -d ./pkg/api
    ```

Swagger documentation is available at http://localhost:3333/swagger/index.html

## E2E Tests

To run E2E tests, run below commands:

    ```bash
         docker-compose build -t
         ./seed_e2e_data.sh
         go test -v ./test
         docker-compose down
    ```

E2E tests are in `test` directory.  
The `./seed_e2e_data.sh` sets up the data needed for E2E tests.  
Once tests are done, the containers are stopped and removed.

There are many other ways to do E2E tests.  
A separate `docker-compose.yml` can be set up which runs the seeding scripts when Postgres service is brought up.  
Depending on the infra and schema setup choices, it can be evolved.