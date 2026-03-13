// Package locationdetails provides location details transformation functionality that extracts
// and structures location information from Opencage geocoding API responses including
// stadium name, city, state, and geographic coordinates.
package locationdetails

import (
	"have-a-nice-pickem-etl/internal/extract/location"
	"time"
)

type New struct {
	location.Location
}

type LocationDetails struct {
	LocationID string    `gorm:"column:id"`
	Stadium    string    `gorm:"column:stadium"`
	City       string    `gorm:"column:city"`
	State      string    `gorm:"column:state"`
	Latitude   float64   `gorm:"column:latitude"`
	Longitude  float64   `gorm:"column:longitude"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (LocationDetails) TableName() string {
	return "pickem.locations"
}

// parseLocationID returns the location ID (maidenhead) from the first Opencage result.
func (l New) parseLocationID() string {
	var locationID string = l.LocationID
	return locationID
}

// parseStadium returns the stadium name from the first Opencage result components.
func (l New) parseStadium() string {
	var stadium string = l.Stadium
	return stadium
}

// parseCity returns the city from the first Opencage result components.
func (l New) parseCity() string {
	var city string = l.City
	return city
}

// parseState returns the state from the first Opencage result components.
func (l New) parseState() string {
	var state string = l.State
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
func (l New) Instantiate() LocationDetails {
	return LocationDetails{
		LocationID: l.parseLocationID(),
		Stadium:    l.parseStadium(),
		City:       l.parseCity(),
		State:      l.parseState(),
		Latitude:   l.parseLatitude(),
		Longitude:  l.parseLongitude(),
	}
}
