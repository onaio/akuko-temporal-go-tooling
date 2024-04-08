package main

import (
	"fmt"
	"os"

	"go.temporal.io/sdk/client"
)

func main() {
	clientOptions := client.Options{
		HostPort:  os.Getenv("TEMPORAL_HOST"),
		Namespace: os.Getenv("SOURCE_CREATION_AND_UPDATING_TEMPORAL_NAMESPACE"),
	}
	temporalClient, err := client.Dial(clientOptions)
	if err != nil {
		panic(err)
	}
	fmt.Println("Health check passed.")
	defer temporalClient.Close()
}
