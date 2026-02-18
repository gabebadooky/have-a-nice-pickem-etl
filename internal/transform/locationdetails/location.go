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

// InstantiateLocationDetails runs the given instantiator and returns the location details.
func InstantiateLocationDetails(i Instantiator) LocationDetails {
	return i.instantiate()
}

// parseLocationID returns the location ID (maidenhead) from the first Opencage result.
func (l New) parseLocationID() string {
	var locationID string = l.Opencage.Results[0].Annotations.Maidenhead
	return locationID
}

// parseStadium returns the stadium name from the first Opencage result components.
func (l New) parseStadium() string {
	var stadium string = l.Opencage.Results[0].Components.Stadium
	return stadium
}

// parseCity returns the city from the first Opencage result components.
func (l New) parseCity() string {
	var city string = l.Opencage.Results[0].Components.City
	return city
}

// parseState returns the state from the first Opencage result components.
func (l New) parseState() string {
	var state string = l.Opencage.Results[0].Components.State
	return state
}

// parseLatitude returns the latitude from the first Opencage result geometry.
func (l New) parseLatitude() float64 {
	var latitude float64 = l.Opencage.Results[0].Geometry.Lat
	return latitude
}

// parseLongitude returns the longitude from the first Opencage result geometry.
func (l New) parseLongitude() float64 {
	var longitude float64 = l.Opencage.Results[0].Geometry.Lon
	return longitude
}

// instantiate builds LocationDetails from the Opencage geocode response.
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
