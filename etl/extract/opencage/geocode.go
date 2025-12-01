package opencage

import (
	"encoding/json"
	"fmt"
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/utils"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Concatenate query string onto Opencage Forward Geocode Endpoint URL
func formatURLwithQueryString(stadium string, city string, state string) string {
	godotenv.Load()
	var apikey string = os.Getenv("OPENCAGE_API_KEY")
	var url string = fmt.Sprintf("%s?key=%s", utils.OPENCAGE_GEOCODE_ENDPOINT_URL, apikey)
	var formattedStadium string = utils.FormatStringID(stadium)
	var formattedCity string = utils.FormatStringID(city)
	var formattedState string = utils.FormatStringID(state)

	if state == "" {
		return fmt.Sprintf("%s&q=%s+%s", url, formattedStadium, formattedCity)
	} else {
		return fmt.Sprintf("%s&q=%s+%s+%s", url, formattedStadium, formattedCity, formattedState)
	}
}

// Call Opencage Forward Geocode API for given stadium, city, state and country
func GetGeocode(stadium string, city string, state string, country string) pickemstructs.OpencageResponse {
	var opencageEndpoint string = formatURLwithQueryString(stadium, city, state)
	log.Printf("\nCalling Opencage API endpoint for %s %s, %s: %s\n", stadium, city, state, opencageEndpoint)
	resp, err := http.Get(opencageEndpoint)
	if err != nil {
		log.Panicf("Error occurred calling Opencage API Endpoint: %s: \n%s\n", opencageEndpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Panicf("Non 200 response code returned from %s:\n%d", opencageEndpoint, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("Error occurred parsing Opencage API Endpoint Response for %s %s, %s: \n%s\n", stadium, city, state, err)
	}

	var geocodeDetails pickemstructs.OpencageResponse
	jsonerr := json.Unmarshal([]byte(body), &geocodeDetails)
	if jsonerr != nil {
		log.Panicf("Error occurred decoding Opencage API Endpoint Response for %s %s, %s: \n%s\n", stadium, city, state, jsonerr)
	}

	//log.Printf("scheduleDetails:\n%v\n", geocodeDetails)
	return geocodeDetails
}
