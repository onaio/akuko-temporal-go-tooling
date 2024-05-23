# akuko-temporal-go-tooling

This Go module provides various Akuko temporal activities written using the golang language. At the moment the activities include functionality for converting a GeoParquet file to GeoJSON and sanitizing GeoJSON feature property names.

## Features

- Convert GeoParquet files to GeoJSON format.
- Sanitize GeoJSON feature property names to [cube dimension naming syntax](https://cube.dev/docs/product/data-modeling/syntax#naming).

## Requirements

- Go version 1.20 or higher.
- Temporal server setup. Refer to [Temporal's documentation](https://docs.temporal.io/dev-guide/go/project-setup#local-dev-server) for instructions on setting up a local dev server.

## Installation

To build the binary, run the following command:

Build

```bash
go build -o ./bin/akuko-temporal-go-tooling
```

Run

```bash
./bin/akuko-temporal-go-tooling
```
or run the go program directly:

Install dependencies
```bash
go mod tidy
```

Run

```bash
gow run main.go
```

## Usage

### Environment Variables

Before running the tool, make sure to set the following environment variables:

- `TEMPORAL_HOST`: The host address of the Temporal server.
- `TEMPORAL_NAMESPACE`: The namespace used for performing this module activities in Temporal.
- `GEOPARQUET_WORKER_TASK_QUEUE_NAME`: The name of the task queue for GeoParquet worker tasks.

### Docker

Building for arm64 based architectures:
```
docker buildx build -t <docker_tag_name> --platform linux/arm64 .
```

Building for amd64 based architectures:
```
docker buildx build -t <docker_tag_name> --platform linux/amd64 .
```

Building for both amd64 and arm64 based architectures:
```
docker buildx build -t <docker_tag_name> --platform linux/amd64, linux/arm64 .
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.
