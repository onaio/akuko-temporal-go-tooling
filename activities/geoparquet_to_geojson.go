package activities

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/apache/arrow/go/v14/parquet/file"
	"github.com/onaio/akuko-geoparquet-temporal-tooling/internal/geojson"
	"github.com/onaio/akuko-geoparquet-temporal-tooling/internal/geoparquet"
)

// ReadFileBytes reads a file from a filepath and returns its contents as a byte slice
func ReadFileBytes(filePath string) ([]byte, error) {
	// Read the file contents
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

// WriteFileBytes writes a byte slice to a file at the given filepath
func WriteFileBytes(filePath string, data []byte) error {
	// Write the data to the file
	err := ioutil.WriteFile(filePath, data, 0644) // 0644 is the file permission
	if err != nil {
		return err
	}

	return nil
}

func newFileReader(filepath string) (*file.Reader, error) {
	f, fileErr := os.Open(filepath)
	if fileErr != nil {
		return nil, fileErr
	}
	return file.NewParquetReader(f)
}

type ConvertGeoParquetToGeoJSONActivityParams struct {
	GeoParquetFilePath string
	GeoJSONFilePath    string
}

type ConvertGeoParquetToGeoJSONActivityReturnType struct {
	FilePath string
	Metadata *geoparquet.Metadata
}

func ConvertGeoParquetToGeoJSONActivity(ctx context.Context, params *ConvertGeoParquetToGeoJSONActivityParams) (*ConvertGeoParquetToGeoJSONActivityReturnType, error) {
	fileBytes, err := ReadFileBytes(params.GeoParquetFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	// Convert from GeoParquet to GeoJSON
	geoJSONBuffer := &bytes.Buffer{}
	err = geojson.FromParquet(bytes.NewReader(fileBytes), geoJSONBuffer)
	if err != nil {
		fmt.Println("Error converting from Parquet to GeoJSON: %v", err)
		return nil, err
	}

	reader, readerErr := newFileReader(params.GeoParquetFilePath)
	if readerErr != nil {
		fmt.Println("Error converting from Parquet to GeoJSON: %v", readerErr)
		return nil, readerErr
	}
	defer reader.Close()

	metadata, metadataErr := geoparquet.GetMetadata(reader.MetaData().GetKeyValueMetadata())
	if metadataErr != nil {
		fmt.Println("Error converting from Parquet to GeoJSON: %v", metadataErr)
		return nil, metadataErr
	}

	fmt.Println("MetaData: %s", &metadata)

	err = WriteFileBytes(params.GeoJSONFilePath, geoJSONBuffer.Bytes())
	if err != nil {
		fmt.Println("Error writing file to disk: %v", err)
		return nil, err
	}
	var data = ConvertGeoParquetToGeoJSONActivityReturnType{
		FilePath: params.GeoJSONFilePath,
		Metadata: metadata,
	}
	return &data, nil
}
