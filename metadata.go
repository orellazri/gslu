package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Metadata struct {
	SourcePath string
}

const METADATA_FILENAME string = "gslu.json"

func CreateMetadataFile(dstDir string, metadata Metadata) error {
	metadataFile, err := os.OpenFile(fmt.Sprintf("%s/%s", dstDir, METADATA_FILENAME), os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer metadataFile.Close()

	jsonBytes, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	metadataFile.Write(jsonBytes)

	return nil
}

func ReadMetadataFile(dir string) (*Metadata, error) {
	fileBytes, err := os.ReadFile(fmt.Sprintf("%s/%s", dir, METADATA_FILENAME))
	if err != nil {
		return nil, err
	}

	var metadata Metadata
	json.Unmarshal(fileBytes, &metadata)

	return &metadata, nil
}
