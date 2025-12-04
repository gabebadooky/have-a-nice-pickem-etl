package opencage

import (
	"encoding/json"
	"fmt"
	"have-a-nice-pickem-etl/etl/utils"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type OpencageResponse struct {
	Results []result `json:"results"`
}

type result struct {
	Geometry geometry `json:"geometry"`
}

type geometry struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lng"`
}

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

// Call Opencage Forward Geocode API and read endpoint response
func callEndpoint(opencageEndpoint string) ([]byte, error) {
	resp, err := http.Get(opencageEndpoint)
	if err != nil {
		return nil, fmt.Errorf("Error occurred calling Opencage API Endpoint: %s: \n%s\n", opencageEndpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("non 200 response code returned from %s:\n%d", opencageEndpoint, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error occurred parsing Opencage API Endpoint Response for %s: \n%s", opencageEndpoint, err)
	}

	return body, nil
}

func decodeResponse(body []byte) (OpencageResponse, error) {
	var geocodeDetails OpencageResponse

	err := json.Unmarshal([]byte(body), &geocodeDetails)
	if err != nil {
		return OpencageResponse{}, fmt.Errorf("error occurred decoding Opencage API Endpoint Response: \n%s", err)
	}

	return geocodeDetails, nil
}

// Retreive Opencage Forward Geocode API Response for given stadium, city, state and country
func GetGeocode(stadium string, city string, state string, country string) OpencageResponse {
	var opencageEndpoint string = formatURLwithQueryString(stadium, city, state)
	log.Printf("\nCalling Opencage API endpoint for %s %s, %s: %s\n", stadium, city, state, opencageEndpoint)

	body, err := callEndpoint(opencageEndpoint)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	geocodeDetails, err := decodeResponse(body)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	//log.Printf("scheduleDetails:\n%v\n", geocodeDetails)
	return geocodeDetails
}
