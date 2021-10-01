package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Metadata struct {
	SourcePath string
}

func CreateMetadataFile(dstDir string, metadata Metadata) error {
	metadataFile, err := os.OpenFile(fmt.Sprintf("%s/gslu.json", dstDir), os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, os.ModePerm)
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
