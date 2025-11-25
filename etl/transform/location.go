package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/location"
)

type Location struct {
	Stadium   string
	City      string
	State     string
	Latitude  float64
	Longitude float64
}

// Transforms and consolidates Location properties from various sources
func CreateLocationRecord(espnGameDetails pickemstructs.ESPNGameDetailsResponse, geocodeDetails pickemstructs.OpencageResponse) Location {
	var newRecord Location

	newRecord.Stadium = location.ParseStadium(espnGameDetails)
	newRecord.City = location.ParseCity(espnGameDetails)
	newRecord.State = location.ParseState(espnGameDetails)
	newRecord.Latitude = location.ParseLatitude(geocodeDetails)
	newRecord.Longitude = location.ParseLongitude(geocodeDetails)

	return newRecord
}
