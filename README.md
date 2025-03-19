# PracGo

Este proyecto consiste en un sistema de recopilación de métricas de una base de datos para monitorizar los datos. Desarrollado como parte de mis prácticas de primer año en el Grado Superior de Desarrollo de Aplicaciones Web en la empresa [iAR Soft](https://www.iar-soft.com/).

# Stack

- Go
- Docker
- MongoDB

# Uso

Crear los contenedores
```
docker compose up --build -d
```

Conectarse al contenedor
```
docker exec -it pracgo bash
```

Ejecutar el programa

```
go run ./cmd/pracgo
```

# Referencias

- https://docs.docker.com/reference/dockerfile/
- https://go.dev/doc/effective_go
- https://gobyexample.com/
- https://www.geeksforgeeks.org/how-to-use-go-with-mongodb/
- https://docs.docker.com/reference/compose-file/
- https://www.mongodb.com/resources/products/fundamentals/create-database
