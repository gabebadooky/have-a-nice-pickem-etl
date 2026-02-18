// Package opencagelocation provides Opencage geocoding API client functionality.
// It calls the Opencage Forward Geocoding API to convert stadium addresses (stadium name,
// city, state) into geographic coordinates and standardized location data.
package opencagelocation

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type OpencageLocation interface {
	geocodeDetails() OpencageEndpoint
}

type OpencageForwardGeocode struct {
	Stadium string
	City    string
	State   string
}

// GetLocationDetails runs the given Opencage location instantiator and returns geocoded location data.
func GetLocationDetails(l OpencageLocation) OpencageEndpoint {
	return l.geocodeDetails()
}

// formatURLwithQueryString builds the Opencage forward geocode URL with API key and address query.
func formatURLwithQueryString(stadium string, city string, state string) string {
	godotenv.Load()
	var apikey string = os.Getenv("OPENCAGE_API_KEY")
	url := fmt.Sprintf("%s?key=%s", utils.OPENCAGE_GEOCODE_ENDPOINT_URL, apikey)
	var formattedStadium string = utils.FormatStringID(stadium)
	var formattedCity string = utils.FormatStringID(city)
	var formattedState string = utils.FormatStringID(state)

	if state == "" {
		queryString := fmt.Sprintf("%s&q=%s+%s", url, formattedStadium, formattedCity)
		return queryString
	} else {
		queryString := fmt.Sprintf("%s&q=%s+%s+%s", url, formattedStadium, formattedCity, formattedState)
		return queryString
	}
}

// decodeOpencageResponse unmarshals the Opencage API response body into OpencageEndpoint.
func decodeOpencageResponse(body []byte) (OpencageEndpoint, error) {
	return utils.DecodeJSON[OpencageEndpoint](body)
}

// geocodeDetails calls the Opencage forward geocode API for the configured stadium, city, and state.
func (g OpencageForwardGeocode) geocodeDetails() OpencageEndpoint {
	opencageEndpoint := formatURLwithQueryString(g.Stadium, g.City, g.State)
	log.Printf("\nCalling Opencage API endpoint for %s %s, %s: %s\n", g.Stadium, g.City, g.State, opencageEndpoint)

	body, err := utils.CallEndpoint(opencageEndpoint)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	geocodeDetails, err := decodeOpencageResponse(body)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	//log.Printf("scheduleDetails:\n%v\n", geocodeDetails)
	return geocodeDetails
}
