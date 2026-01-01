package location

import opencagelocation "have-a-nice-pickem-etl/etl/extract/location/opencage"

type AllLocationInfo interface {
	locationInfo() Location
}

type OpencageLocationInfo struct {
	Stadium string
	City    string
	State   string
}

type Location struct {
	Opencage opencagelocation.OpencageEndpoint
}

func ConsolidateLocationInfo(l AllLocationInfo) Location {
	return l.locationInfo()
}

func (l OpencageLocationInfo) locationInfo() Location {
	opencageLocation := opencagelocation.GetLocationDetails(
		opencagelocation.OpencageForwardGeocode{
			Stadium: l.Stadium,
			City:    l.City,
			State:   l.State,
		},
	)

	return Location{
		Opencage: opencageLocation,
	}
}
