package openweathermap

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type OpenWeatherTimestampForecast struct {
	Lat          float64
	Lon          float64
	ZuluGameTime string
}

func formatURLwithQueryString(lat float64, lon float64, zuluGameTime string) string {
	godotenv.Load()
	var apikey string = os.Getenv("fb4548e7e50ecf86dbf87a952860a254")
	var unixTimestamp string = utils.ConvertZuluTimestampToUnixTime(zuluGameTime)
	url := fmt.Sprintf("%s?lat=%s&lon=%s&dt=%s&appid=%s", lat, lon, unixTimestamp, apikey)
	return url
}

func decodeOpenWeatherResponse(body []byte) (OpenWeatherMapEndpoint, error) {
	return utils.DecodeJSON[OpenWeatherMapEndpoint](body)
}

func (w OpenWeatherTimestampForecast) GetForecastDetails() OpenWeatherMapEndpoint {
	openWeatherEndpoint := formatURLwithQueryString(w.Lat, w.Lon, w.ZuluGameTime)
	log.Printf("\nCalling OpenWeather API endpoint for Lat/Lon: %f/%f on %s", w.Lat, w.Lon, w.ZuluGameTime)

	body, err := utils.CallEndpoint(openWeatherEndpoint)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	forecastDetails, err := decodeOpenWeatherResponse(body)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return forecastDetails
}
