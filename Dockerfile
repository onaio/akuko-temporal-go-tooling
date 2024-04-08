FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go build -o ./bin/akuko-geoparquet-temporal-tooling

ENTRYPOINT ["./bin/akuko-geoparquet-temporal-tooling"]
