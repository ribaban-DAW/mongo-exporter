# mongo-exporter

Este proyecto consiste en un sistema de recolección de métricas de una base de datos para monitorizar los datos. Desarrollado como parte de mis prácticas de primer año en el Grado Superior de Desarrollo de Aplicaciones Web en la empresa [iAR Soft](https://www.iar-soft.com/).

# API Endpoints

|Método|Endpoint|Descripción|
|-|-|-|
|`GET`|`/v1/students`|Obtener todos los estudiantes| 
|`GET`|`/v1/students/:id`|Obtener un estudiante por ID| 

# Instalación

Clonar el repositorio
```
git clone https://github.com/ribaban-DAW/mongo-exporter
```

Navegar al directorio del proyecto
```
cd mongo-exporter
```

Instalar dependencias
```
go mod tidy
```

# Uso

Ejecuta el programa
```
go run ./cmd/mongo-exporter
```

De manera alternativa, puedes crear el ejecutable y ejecutarlo.
```
go build ./cmd/mongo-exporter
./mongo-exporter
```

Una vez esté en ejecución, puedes interactuar con la API utilizando tu navegador, curl o cualquier otro método que prefieras.

- Para obtener una lista de todos los estudiante:
```
curl localhost:8080/v1/students
```

- Para obtener el estudiante por ID:
```
curl localhost:8080/v1/students/1
```

Revisa [API Endpoints](#api-endpoints) para ver los endpoints disponibles.

# Referencias

- https://go.dev/doc/tutorial/web-service-gin
- https://youtu.be/67yGbvyM1is
- https://gin-gonic.com/docs
- https://github.com/gin-gonic/examples/tree/master/group-routes
