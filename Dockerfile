FROM golang:1.20 AS base

WORKDIR /srv/app

COPY . .

FROM base as default
RUN go build -o ./bin/akuko-temporal-go-tooling
WORKDIR /srv/app/healthcheck
RUN go build -o ../bin/healthcheck
WORKDIR /srv/app
ENTRYPOINT ["./bin/akuko-temporal-go-tooling"]
