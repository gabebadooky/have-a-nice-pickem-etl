package transform

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/extract/opencage"
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/location"
	"have-a-nice-pickem-etl/etl/utils"
)

// Transforms and consolidates Location properties from various sources
func CreateLocationRecord(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties) pickemstructs.Location {
	var stadium string = location.ParseStadium(consolidatedGameProperties)
	var city string = location.ParseCity(consolidatedGameProperties)
	var state string = location.ParseState(consolidatedGameProperties)
	var country string = location.ParseCountry(consolidatedGameProperties)
	var geocodeDetails pickemstructs.OpencageResponse = opencage.GetGeocode(stadium, city, state, country)
	var locationID string = fmt.Sprintf("%s-%s-%s", utils.FormatStringID(stadium), utils.FormatStringID(city), utils.FormatStringID(state))

	var newRecord pickemstructs.Location = pickemstructs.Location{
		LocationID: locationID,
		Stadium:    stadium,
		City:       city,
		State:      state,
		Latitude:   location.ParseLatitude(geocodeDetails),
		Longitude:  location.ParseLongitude(geocodeDetails),
	}

	return newRecord
}
