# mongo-exporter

This project consists in a Metric Collection System for a database to monitor the data. It's developped as part of my first year internship in [iAR Soft](https://www.iar-soft.com/).

# API Endpoints

|Method|Endpoint|Description|
|-|-|-|
|`GET`|`/v1/students`|Get all students| 
|`GET`|`/v1/students/:id`|Get student by ID| 

# Installation

Clone the repository
```
git clone https://github.com/ribaban-DAW/mongo-exporter
```

Navigate to the project folder
```
cd mongo-exporter
```

Install dependencies
```
go mod tidy
```

# Usage

Run the program
```
go run ./cmd/mongo-exporter
```

Alternatively, you can build an executable and run it
```
go build ./cmd/mongo-exporter
./mongo-exporter
```

Once it's running, you can interact with the API using your browser, curl, or any method you prefer.
- To get a list of every student:
```
curl localhost:8080/v1/students
```
- To get the student by ID:
```
curl localhost:8080/v1/students/1
```

Check [API Endpoints](#api-endpoints) to see available endpoints.

# References

- https://go.dev/doc/tutorial/web-service-gin
- https://youtu.be/67yGbvyM1is
- https://gin-gonic.com/docs
- https://github.com/gin-gonic/examples/tree/master/group-routes
