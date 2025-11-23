package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/location"
)

// Transforms and consolidates Game properties from various sources
func CreateLocationRecord(espnGameDetails pickemstructs.ESPNGameDetailsResponse, geocodeDetails pickemstructs.OpencageResponse) pickemstructs.Location {
	var newRecord pickemstructs.Location

	newRecord.Stadium = location.ParseStadium(espnGameDetails)
	newRecord.City = location.ParseCity(espnGameDetails)
	newRecord.State = location.ParseState(espnGameDetails)
	newRecord.Latitude = location.ParseLatitude(geocodeDetails)
	newRecord.Longitude = location.ParseLongitude(geocodeDetails)

	return newRecord
}
