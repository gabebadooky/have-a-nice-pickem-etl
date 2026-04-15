package forecast

import "have-a-nice-pickem-etl/internal/extract/forecast/openweathermap"

type OpenWeatherForecast struct {
	LocationID   string
	ZuluGameTime string
	Lat          float64
	Lon          float64
}

type Forecast struct {
	LocationID   string
	ZuluGameTime string
	OpenWeather  openweathermap.OpenWeatherMapEndpoint
}

func (w OpenWeatherForecast) GetForecast() Forecast {
	openWeatherTimestampForecast := openweathermap.OpenWeatherTimestampForecast{
		Lat:          w.Lat,
		Lon:          w.Lon,
		ZuluGameTime: w.ZuluGameTime,
	}

	openWeatherForecastDetails := openWeatherTimestampForecast.GetForecastDetails()

	return Forecast{
		LocationID:   w.LocationID,
		ZuluGameTime: w.ZuluGameTime,
		OpenWeather:  openWeatherForecastDetails,
	}
}
