**Sensor API**

This repository is backend implementation of Sensor API which collects Sensor data and calculates various aggregation statistics. 
The API is built using the Go language and the Gin framework.

**Features**

* Captures Sensor group and Sensors information
* Collects data from sensors
* Performs aggregations on sensor data

**Tech Stack**

* GoLang
* PostgresSQL for data storage
* Docker 
* Swagger API documentation
* Redis

**Prerequisites**

Ensure the below are installed:
- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)

**Local Setup:**

1. **Clone the repository:**

    ```bash
    git clone https://github.com/hriteshg/sensor-api.git
    cd sensor-api
    ```

2. **Rename `env.example` to  `.env` file in the project root and configure the data**

**Running the API**

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


**Swagger Documentation**

Swagger documentation is available at http://localhost:3333/swagger/index.html