package load

import (
	"fmt"
	"os"
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
	folderName := fmt.Sprintf("%d%02d%d%d%d", year, month, day, hour, min)
	return folderName
}

// InstantiateLoadDirectory concatenates the folder path for the data load path
func InstantiateLoadDirectory() string {
	newFolderName := instantiateDestinationFolderName()
	directoryPathStr := fmt.Sprintf("./data/%s", newFolderName)
	os.Mkdir(directoryPathStr, 0777)
	return directoryPathStr
}
