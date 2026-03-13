// Package location provides location extraction functionality that uses the Opencage
// geocoding API to convert stadium addresses into geographic coordinates and
// standardized location data.
package location

import opencagelocation "have-a-nice-pickem-etl/internal/extract/location/opencage"

type OpencageLocation struct {
	LocationID string
	Stadium    string
	City       string
	State      string
}

type Location struct {
	LocationID string
	Stadium    string
	City       string
	State      string
	Opencage   opencagelocation.OpencageEndpoint
}

// extractLocation geocodes the venue via Opencage and returns location details.
func (l OpencageLocation) GetLocation() Location {
	opencageForwardGeocode := opencagelocation.OpencageForwardGeocode{
		Stadium: l.Stadium,
		City:    l.City,
		State:   l.State,
	}

	opencageLocationDetails := opencageForwardGeocode.GetGeocodeDetails()

	return Location{
		LocationID: l.LocationID,
		Stadium:    l.Stadium,
		City:       l.City,
		State:      l.State,
		Opencage:   opencageLocationDetails,
	}
}
