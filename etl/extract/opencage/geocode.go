package opencage

import (
	"encoding/json"
	"fmt"
	"have-a-nice-pickem-etl/etl/extract"
	"have-a-nice-pickem-etl/etl/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Geocode struct {
	stadium string
	city    string
	state   string
}

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

func decodeOpencageResponse(body []byte) (OpencageResponse, error) {
	var geocodeDetails OpencageResponse

	err := json.Unmarshal([]byte(body), &geocodeDetails)
	if err != nil {
		return OpencageResponse{}, fmt.Errorf("error occurred decoding Opencage API Endpoint Response: \n%s", err)
	}

	return geocodeDetails, nil
}

// Retreive Opencage Forward Geocode API Response for given stadium, city, state and country
func (g Geocode) Get() OpencageResponse {
	var opencageEndpoint string = formatURLwithQueryString(g.stadium, g.city, g.state)
	log.Printf("\nCalling Opencage API endpoint for %s %s, %s: %s\n", g.stadium, g.city, g.state, opencageEndpoint)

	body, err := extract.CallEndpoint(opencageEndpoint)
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
