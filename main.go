package main

import (
	"fmt"
	"log"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/onaio/akuko-temporal-go-tooling/activities"
	"github.com/onaio/akuko-temporal-go-tooling/workflows"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	clientOptions := client.Options{
		HostPort:  os.Getenv("TEMPORAL_HOST"),
		Namespace: os.Getenv("TEMPORAL_NAMESPACE"),
	}
	temporalClient, err := client.Dial(clientOptions)
	if err != nil {
		fmt.Println("Unable to create a Temporal Client: %s", err)
	}
	if temporalClient != nil {
		defer temporalClient.Close()
		// Create a new Worker
		yourWorker := worker.New(temporalClient, os.Getenv("GEOPARQUET_WORKER_TASK_QUEUE_NAME"), worker.Options{})
		// Register Workflows
		yourWorker.RegisterWorkflow(workflows.ConvertGeoparquetToGeojson)
		// Register Activities
		yourWorker.RegisterActivity(activities.ConvertGeoParquetToGeoJSONActivity)
		yourWorker.RegisterActivity(activities.SanitizeGeoJSONFeaturePropertiesActivity)
		// Start the Worker Process
		err = yourWorker.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalln("Unable to start the Worker Process", err)
		}
	}
}
