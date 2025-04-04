# mongo-exporter

Este proyecto consiste en un sistema de recolección de métricas de una base de
datos para monitorizar los datos. Desarrollado como parte de mis prácticas de
primer año en el Grado Superior de Desarrollo de Aplicaciones Web en la empresa
[iAR Soft](https://www.iar-soft.com/).

# API Endpoints

| Método | Endpoint                  | Descripción                                                                                              |
| ------ | ------------------------- | -------------------------------------------------------------------------------------------------------- |
| `GET`  | `/v1/healthcheck`         | Obtener el estado de la API                                                                              |
| `GET`  | `/v1/hello`               | Obtener el mensaje "hello world"                                                                         |
| `GET`  | `/v1/metrics/collection`  | Obtener métricas relacionadas a una colección. Requiere los parámetros de consulta `dbName` y `collName` |
| `GET`  | `/v1/metrics/connections` | Obtener métricas relacionadas con conexiones del servidor                                                |
| `GET`  | `/v1/metrics/cpu`         | Obtener métricas relacionadas con CPU del servidor                                                       |
| `GET`  | `/v1/metrics/opcounters`  | Obtener métricas relacionadas con operaciones del servidor                                               |
| `GET`  | `/v1/metrics/ram`         | Obtener métricas relacionadas con RAM del servidor                                                       |

# Uso

Clona el repositorio:

```
git clone https://github.com/SrVariable/mongo-exporter
```

Navega al directorio del proyecto:

```
cd mongo-exporter
```

Crea el archivo `.env` siguiendo el `.env.example` para configurar las
variables de entorno. Para la configuración por defecto, copia `.env.example` a
`.env`:

```
cp .env.example .env
```

El `.env` debería ser similar a esto:

```
APP_PORT=8080

DB_NAME=MyDatabaseName
DB_HOST=mongo
DB_PORT=27017

GRAFANA_USER=admin
GRAFANA_PASS=grafana
```

Construye los contenedores:

```
make
```

> [!NOTE]
>
> Si no tienes `make`, puedes ejecutar:
>
> ```
> docker compose down
> docker compose up --build -d
> ```

Una vez esté construido, puedes interactuar con la API utilizando tu navegador,
curl o cualquier otro método que prefieras.

- Para obtener métricas relacionadas con CPU

```
curl http://localhost:8080/v1/metrics/cpu
```

- Para obtener métricas relacionadas a la Colección `bar` de la Base de datos `foo`:

```
curl http://localhost:8080/v1/metrics/collection?dbName=foo&collName=bar
```

Revisa [API Endpoints](#api-endpoints) para ver los endpoints disponibles.

# Referencias

- https://go.dev/doc/tutorial/web-service-gin
- https://youtu.be/67yGbvyM1is
- https://gin-gonic.com/docs
- https://github.com/gin-gonic/examples/tree/master/group-routes
- https://stackoverflow.com/questions/33322103/multiple-froms-what-it-means
- https://stackoverflow.com/questions/75973805/creating-dockerfile-for-golang-web-application
- https://www.docker.com/blog/developing-go-apps-docker/
- https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo
- https://www.youtube.com/watch?v=bDWApqAUjEI
- https://www.youtube.com/watch?v=g7cNQB2kCgE
- https://www.mongodb.com/docs/manual/reference/command/serverStatus/
- https://github.com/docker/awesome-compose/tree/master/prometheus-grafana
