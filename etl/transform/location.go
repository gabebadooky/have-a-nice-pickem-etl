package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/location"
)

// Transforms and consolidates Location properties from various sources
func CreateLocationRecord(espnGameDetails pickemstructs.ESPNGameDetailsResponse, geocodeDetails pickemstructs.OpencageResponse) pickemstructs.Location {
	var newRecord pickemstructs.Location = pickemstructs.Location{
		Stadium:   location.ParseStadium(espnGameDetails),
		City:      location.ParseCity(espnGameDetails),
		State:     location.ParseState(espnGameDetails),
		Latitude:  location.ParseLatitude(geocodeDetails),
		Longitude: location.ParseLongitude(geocodeDetails),
	}

	return newRecord
}
