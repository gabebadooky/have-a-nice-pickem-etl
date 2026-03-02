package load

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

// instantiateCsvWriter creates and returns a CSV writer for the given file path.
func instantiateCsvWriter(filepath string) (*os.File, *csv.Writer) {
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Print(err)
		log.Fatalf("Error occurred instantiating CSV Writer to %s", filepath)
	}

	csvwriter := csv.NewWriter(f)
	return f, csvwriter
}

// instantiateDestinationFolderName returns a timestamp-based folder name for archiving (e.g. 2026-2-17-12:30).
func instantiateDestinationFolderName() string {
	timestamp := time.Now().Local().UTC()
	year := timestamp.Year()
	month := timestamp.Month()
	day := timestamp.Day()
	hour := timestamp.Hour()
	min := timestamp.Minute()
	folderName := fmt.Sprintf("%d%02d%d%d%d", year, month, day, hour, min)
	return folderName
}

// InstantiateLoadDirectory concatenates the folder path for the data load path
func instantiateLoadDirectory() string {
	newFolderName := instantiateDestinationFolderName()
	directoryPathStr := fmt.Sprintf("./data/%s", newFolderName)
	os.Mkdir(directoryPathStr, 0777)
	return directoryPathStr
}
