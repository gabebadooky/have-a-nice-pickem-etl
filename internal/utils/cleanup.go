package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

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
