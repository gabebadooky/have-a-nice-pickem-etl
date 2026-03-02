package load

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/transform/locationdetails"
	"log"
)

// LocationDetails writes location detail records to data/locations.csv.
func loadLocationDetails(records []locationdetails.LocationDetails, csvLoadFolderPath string) {
	bulkDataLoadFilePath := fmt.Sprintf("%s/%s", csvLoadFolderPath, "locations.csv")
	f, w := instantiateCsvWriter(bulkDataLoadFilePath)
	defer f.Close()
	defer w.Flush()

	log.Printf("Writing Location Details records to %s", bulkDataLoadFilePath)

	for _, record := range records {
		w.Write([]string{
			record.LocationID,
			record.Stadium,
			record.City,
			record.State,
			fmt.Sprintf("%f", record.Latitude),
			fmt.Sprintf("%f", record.Longitude),
		})
	}

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	queryString := fmt.Sprintf("CALL %s('%s')", "bulk_load_stats", bulkDataLoadFilePath)
	callBulkLoadProcedure(queryString)
}
