package location

import "have-a-nice-pickem-etl/etl/pickemstructs"

func ParseLatitude(geocodeDetails pickemstructs.OpencageResponse) float64 {
	var latitude float64 = geocodeDetails.Results[0].Geometry.Lat
	return latitude
}

func ParseLongitude(geocodeDetails pickemstructs.OpencageResponse) float64 {
	var longitude float64 = geocodeDetails.Results[0].Geometry.Lon
	return longitude
}
