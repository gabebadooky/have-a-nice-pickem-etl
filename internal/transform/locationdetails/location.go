// Package locationdetails provides location details transformation functionality that extracts
// and structures location information from Opencage geocoding API responses including
// stadium name, city, state, and geographic coordinates.
package locationdetails

import "have-a-nice-pickem-etl/internal/extract/location"

type Instantiator interface {
	instantiate() LocationDetails
}

type New struct {
	location.Location
}

type LocationDetails struct {
	LocationID string
	Stadium    string
	City       string
	State      string
	Latitude   float64
	Longitude  float64
}

func InstantiateLocationDetails(i Instantiator) LocationDetails {
	return i.instantiate()
}

func (l New) parseLocationID() string {
	var locationID string = l.Opencage.Results[0].Annotations.Maidenhead
	return locationID
}

func (l New) parseStadium() string {
	var stadium string = l.Opencage.Results[0].Components.Stadium
	return stadium
}

func (l New) parseCity() string {
	var city string = l.Opencage.Results[0].Components.City
	return city
}

func (l New) parseState() string {
	var state string = l.Opencage.Results[0].Components.State
	return state
}

func (l New) parseLatitude() float64 {
	var latitude float64 = l.Opencage.Results[0].Geometry.Lat
	return latitude
}

func (l New) parseLongitude() float64 {
	var longitude float64 = l.Opencage.Results[0].Geometry.Lon
	return longitude
}

func (l New) instantiate() LocationDetails {
	return LocationDetails{
		LocationID: l.parseLocationID(),
		Stadium:    l.parseStadium(),
		City:       l.parseCity(),
		State:      l.parseState(),
		Latitude:   l.parseLatitude(),
		Longitude:  l.parseLongitude(),
	}
}
