package location

import opencagelocation "have-a-nice-pickem-etl/etl/extract/location/opencage"

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

func ConsolidateLocationInfo(l locationInstantiator) Location {
	return l.extractLocation()
}

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
