package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"go.temporal.io/sdk/activity"
)

// sanitizeStringToCubeSyntax sanitizes a string to conform to Cube syntax rules.
func sanitizeStringToCubeSyntax(str string) string {
	// Cube doesn't allow numbers as the first character.
	firstCharacter := str[:1]
	if strings.Contains("0123456789", firstCharacter) {
		str = "_" + str
	}

	// Replace non-alphanumeric characters with underscore and convert to lowercase.
	reg := regexp.MustCompile("[^A-Za-z0-9]+")
	sanitized := reg.ReplaceAllString(str, "_")
	return strings.ToLower(sanitized)
}

// sanitizeProperties sanitizes the keys in the properties map to conform to Cube syntax rules.
func sanitizeProperties(properties map[string]interface{}, feature_id int) map[string]interface{} {
	sanitizedProperties := make(map[string]interface{})

	for key, value := range properties {
		sanitizedKey := sanitizeStringToCubeSyntax(key)
		sanitizedProperties[sanitizedKey] = value
		sanitizedProperties["feature_id"] = feature_id
	}

	return sanitizedProperties
}

// SanitizeGeoJSONFile reads a GeoJSON file, sanitizes the feature property names, and writes the sanitized GeoJSON to a new file.
func SanitizeGeoJSONFile(inputFile string, outputFile string) error {
	// Read input GeoJSON file
	geoJSONData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("error reading input GeoJSON file: %v", err)
	}

	// Unmarshal GeoJSON data
	var geoJSONMap map[string]interface{}
	if err := json.Unmarshal(geoJSONData, &geoJSONMap); err != nil {
		return fmt.Errorf("error unmarshalling GeoJSON data: %v", err)
	}

	// Extract features
	features, ok := geoJSONMap["features"].([]interface{})
	if !ok {
		return fmt.Errorf("no features found in GeoJSON data")
	}

	feature_id := 0
	// Sanitize feature properties
	for _, feature := range features {
		feature_id++
		featureMap, ok := feature.(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid feature format")
		}
		properties, ok := featureMap["properties"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("no properties found in feature")
		}
		// sanitizeProperties sanitizes the keys in the properties map to conform to Cube syntax rules.
		// It also adds a feature_id key for each feature with value index + 1.
		featureMap["properties"] = sanitizeProperties(properties, feature_id)
	}

	// Marshal sanitized GeoJSON data
	sanitizedGeoJSONData, err := json.Marshal(geoJSONMap)
	if err != nil {
		return fmt.Errorf("error marshalling sanitized GeoJSON data: %v", err)
	}

	// Write sanitized GeoJSON data to output file
	if err := ioutil.WriteFile(outputFile, sanitizedGeoJSONData, 0644); err != nil {
		return fmt.Errorf("error writing sanitized GeoJSON data to output file: %v", err)
	}

	return nil
}

type SanitizeGeoJSONFeaturePropertiesActivityParams struct {
	GeoJSONFilePath string
}

type SanitizeGeoJSONFeaturePropertiesActivityReturnType struct {
	FilePath string
}

func SanitizeGeoJSONFeaturePropertiesActivity(ctx context.Context, params *SanitizeGeoJSONFeaturePropertiesActivityParams) (*SanitizeGeoJSONFeaturePropertiesActivityReturnType, error) {
	message := "Sanitizing geojson file"
	activity.RecordHeartbeat(ctx, message)
	inputFile := params.GeoJSONFilePath
	outputFile := params.GeoJSONFilePath

	err := SanitizeGeoJSONFile(inputFile, outputFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	message = "GeoJSON file sanitized successfully."
	activity.RecordHeartbeat(ctx, message)
	fmt.Println("GeoJSON file sanitized successfully.")
	var data = SanitizeGeoJSONFeaturePropertiesActivityReturnType{
		FilePath: outputFile,
	}
	return &data, nil
}
