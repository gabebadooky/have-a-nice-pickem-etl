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
	var apikey string = os.Getenv("OPENWEATHER_API_KEY")
	var unixTimestamp string = utils.ConvertZuluTimestampToUnixTime(zuluGameTime)
	url := fmt.Sprintf("%s?lat=%f&lon=%f&dt=%s&appid=%s", utils.OPENWEATHERMAP_TIMESTAMP_ENDPOINT_URL, lat, lon, unixTimestamp, apikey)
	fmt.Printf("\nurl: %s\n", url)
	return url
}

func decodeOpenWeatherResponse(body []byte) (OpenWeatherMapEndpoint, error) {
	return utils.DecodeJSON[OpenWeatherMapEndpoint](body)
}

func (w OpenWeatherTimestampForecast) GetForecastDetails() (OpenWeatherMapEndpoint, error) {
	openWeatherEndpoint := formatURLwithQueryString(w.Lat, w.Lon, w.ZuluGameTime)
	log.Printf("\nCalling OpenWeather API endpoint for Lat/Lon: %f/%f on %s", w.Lat, w.Lon, w.ZuluGameTime)

	body, err := utils.CallEndpoint(openWeatherEndpoint)
	if err != nil {
		return OpenWeatherMapEndpoint{}, fmt.Errorf("Error occured retrieving forecast for (%f, %f) on %s: %s", w.Lat, w.Lon, w.ZuluGameTime, err.Error())
	}

	forecastDetails, err := decodeOpenWeatherResponse(body)
	if err != nil {
		return OpenWeatherMapEndpoint{}, fmt.Errorf("Error occured retrieving forecast for (%f, %f) on %s: %s", w.Lat, w.Lon, w.ZuluGameTime, err.Error())
	}

	return forecastDetails, nil
}
