package load

import (
	"have-a-nice-pickem-etl/internal/transform/locationdetails"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// LocationDetails writes location detail records to data/locations.csv.
func loadLocationDetails(records []locationdetails.LocationDetails, db *gorm.DB) {
	db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"stadium",
			"city",
			"state",
			"latitude",
			"longitude",
			"updated_at",
		}),
	}).Create(&records)

	/*
		db.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"stadium":    gorm.Expr("EXCLUDED.stadium"),
				"city":       gorm.Expr("EXCLUDED.city"),
				"state":      gorm.Expr("EXCLUDED.state"),
				"latitude":   gorm.Expr("EXCLUDED.latitude"),
				"longitude":  gorm.Expr("EXCLUDED.longitude"),
				"updated_at": gorm.Expr("CURRENT_TIMESTAMP"),
			}),
		}).Create(&records)

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
	*/
}
