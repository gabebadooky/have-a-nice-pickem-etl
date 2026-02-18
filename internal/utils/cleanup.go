// Package utils provides shared constants and utility functions used across the ETL pipeline.
// This file handles data archiving and destination folder naming.
package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// instantiateDestinationFolderName returns a timestamp-based folder name for archiving (e.g. 2026-2-17-12:30).
func instantiateDestinationFolderName() string {
	timestamp := time.Now().Local().UTC()
	year := timestamp.Year()
	month := timestamp.Month()
	day := timestamp.Day()
	hour := timestamp.Hour()
	min := timestamp.Minute()
	folderName := fmt.Sprintf("%d-%02d-%d-%d:%d", year, month, day, hour, min)
	return folderName
}

// ArchiveData reads the data directory and prepares to archive existing files (destination folder is computed but not yet applied).
func ArchiveData() {
	files, err := os.ReadDir("./data/")
	if err != nil {
		log.Fatal(err)
	}

	if len(files) > 0 {
		destinationFolder := instantiateDestinationFolderName()
		filepath.Join(destinationFolder, "./data")
		// os.RemoveAll("./data/")
	}
}
