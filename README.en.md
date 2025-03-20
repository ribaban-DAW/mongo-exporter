# mongo-exporter

This project consists in a Metric Collection System for a database to monitor the data. It's developped as part of my first year internship in [iAR Soft](https://www.iar-soft.com/).

# API Endpoints

|Method|Endpoint|Description|
|-|-|-|
|`GET`|`/v1/healthcheck`|Get the health status of the API|
|`GET`|`/v1/hello`|Get "hello world" message|
|`GET`|`/v1/students`|Get a list of all students|
|`GET`|`/v1/students/:id`|Get a student by ID|

# Usage

Clone the repository
```
git clone https://github.com/SrVariable/mongo-exporter
```

Navigate to the project folder
```
cd mongo-exporter
```

Build the containers
```
make
```

> [!NOTE]
>
> If you don't have `make`, you can run:
> ```
> docker compose down
> docker compose up --build -d
> ```

Once it's built, you can interact with the API using your browser, curl, or any method you prefer.
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
- https://stackoverflow.com/questions/33322103/multiple-froms-what-it-means
- https://stackoverflow.com/questions/75973805/creating-dockerfile-for-golang-web-application
- https://www.docker.com/blog/developing-go-apps-docker/
