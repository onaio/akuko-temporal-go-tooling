FROM golang:1.20 AS base

WORKDIR /app

COPY . .

FROM base as production
RUN go build -o ./bin/akuko-temporal-go-tooling
ENTRYPOINT ["./bin/akuko-temporal-go-tooling"]

FROM base as dev
RUN go mod tidy
CMD ["gow", "run", "main.go"]
