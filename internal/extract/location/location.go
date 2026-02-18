// Package location provides location extraction functionality that uses the Opencage
// geocoding API to convert stadium addresses into geographic coordinates and
// standardized location data.
package location

import opencagelocation "have-a-nice-pickem-etl/internal/extract/location/opencage"

type OpencageLocation struct {
	Stadium string
	City    string
	State   string
}

type Location struct {
	Opencage opencagelocation.OpencageEndpoint
}

type locationInstantiator interface {
	extractLocation() Location
}

// ConsolidateLocationInfo runs the given location instantiator and returns the consolidated Location.
func ConsolidateLocationInfo(l locationInstantiator) Location {
	return l.extractLocation()
}

// extractLocation geocodes the venue via Opencage and returns location details.
func (l OpencageLocation) extractLocation() Location {
	var opencageLocationDetails opencagelocation.OpencageEndpoint

	opencageForwardGeocode := opencagelocation.OpencageForwardGeocode{
		Stadium: l.Stadium,
		City:    l.City,
		State:   l.State,
	}

	opencageLocationDetails = opencagelocation.GetLocationDetails(opencageForwardGeocode)

	return Location{
		Opencage: opencageLocationDetails,
	}
}
