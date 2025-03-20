# Build
FROM golang:1.24.1 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /mongo-exporter ./cmd/mongo-exporter

# Runner, for production
FROM alpine:3.21.3 AS run

COPY --from=build /mongo-exporter /mongo-exporter

WORKDIR /app
EXPOSE 8080
CMD ["/mongo-exporter"]
