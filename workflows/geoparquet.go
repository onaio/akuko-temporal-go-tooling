package workflows

import (
	"time"

	"go.temporal.io/sdk/workflow"

	"github.com/onaio/akuko-temporal-go-tooling/activities"
)

func ConvertGeoparquetToGeojson(ctx workflow.Context, params *activities.ConvertGeoParquetToGeoJSONActivityParams) (activities.ConvertGeoParquetToGeoJSONActivityReturnType, error) {
	// Define the Activity Execution options
	// StartToCloseTimeout or ScheduleToCloseTimeout must be set
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	// Execute the Activity synchronously (wait for the result before proceeding)
	var data activities.ConvertGeoParquetToGeoJSONActivityReturnType
	err := workflow.ExecuteActivity(ctx, activities.ConvertGeoParquetToGeoJSONActivity, params).Get(ctx, &data)
	if err != nil {
		return activities.ConvertGeoParquetToGeoJSONActivityReturnType{}, err
	}
	// Make the results of the Workflow available
	return data, nil
}
