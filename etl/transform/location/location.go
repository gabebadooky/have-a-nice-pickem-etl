package location

import "have-a-nice-pickem-etl/etl/extract/location"

type Instantiator interface {
	instantiate() Location
}

type NewLocation struct {
	LocationExtract location.Location
}

type Location struct {
	LocationID string
	Stadium    string
	City       string
	State      string
	Latitude   float64
	Longitude  float64
}

func InstantiateLocation(i Instantiator) Location {
	return i.instantiate()
}

func (l NewLocation) parseLocationID() string {
	var locationID string = l.LocationExtract.Opencage.Results[0].Annotations.Maidenhead
	return locationID
}

func (l NewLocation) parseStadium() string {
	var stadium string = l.LocationExtract.Opencage.Results[0].Components.Stadium
	return stadium
}

func (l NewLocation) parseCity() string {
	var city string = l.LocationExtract.Opencage.Results[0].Components.City
	return city
}

func (l NewLocation) parseState() string {
	var state string = l.LocationExtract.Opencage.Results[0].Components.State
	return state
}

func (l NewLocation) parseLatitude() float64 {
	var latitude float64 = l.LocationExtract.Opencage.Results[0].Geometry.Lat
	return latitude
}

func (l NewLocation) parseLongitude() float64 {
	var longitude float64 = l.LocationExtract.Opencage.Results[0].Geometry.Lon
	return longitude
}

func (l NewLocation) instantiate() Location {
	return Location{
		LocationID: l.parseLocationID(),
		Stadium:    l.parseStadium(),
		City:       l.parseCity(),
		State:      l.parseState(),
		Latitude:   l.parseLatitude(),
		Longitude:  l.parseLongitude(),
	}
}
