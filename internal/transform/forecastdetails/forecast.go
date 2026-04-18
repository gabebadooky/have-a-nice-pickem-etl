package forecastdetails

import (
	"have-a-nice-pickem-etl/internal/extract/forecast"
	"time"
)

type New struct {
	forecast.Forecast
}

type ForecastDetails struct {
	LocationID       string    `gorm:"column:location_id"`
	ZuluTimestamp    string    `gorm:"column:zulu_timestamp"`
	Temperature      float32   `gorm:"column:temperature"`
	FeelsLike        float32   `gorm:"column:feels_like"`
	Humidity         float32   `gorm:"column:humidity"`
	Visibility       float32   `gorm:"column:visibility"`
	WindSpeed        float32   `gorm:"column:wind_speed"`
	ShortDescription string    `gorm:"column:short_description"`
	LongDescription  string    `gorm:"column:long_description"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}

func (f New) parseTemperature() float32 {
	var temperature float32 = f.OpenWeather.Data[0].Temperature
	return temperature
}

func (f New) parseFeelsLike() float32 {
	var feelsLike float32 = f.OpenWeather.Data[0].FeelsLike
	return feelsLike
}

func (f New) parseHumidity() float32 {
	var humidity float32 = f.OpenWeather.Data[0].Humidity
	return humidity
}

func (f New) parseVisibility() float32 {
	var visibility float32 = f.OpenWeather.Data[0].Visibility
	return visibility
}

func (f New) parseWindSpeed() float32 {
	var windSpeed float32 = f.OpenWeather.Data[0].WindSpeed
	return windSpeed
}

func (f New) parseShortDescription() string {
	var shortDescription string = f.OpenWeather.Data[0].Weather[0].ShortDescription
	return shortDescription
}

func (f New) parseLongDescription() string {
	var longDescription string = f.OpenWeather.Data[0].Weather[0].LongDescription
	return longDescription
}

func (f New) Instantiate() ForecastDetails {
	return ForecastDetails{
		LocationID:       f.LocationID,
		ZuluTimestamp:    f.ZuluGameTime,
		Temperature:      f.parseTemperature(),
		FeelsLike:        f.parseFeelsLike(),
		Humidity:         f.parseHumidity(),
		Visibility:       f.parseVisibility(),
		WindSpeed:        f.parseWindSpeed(),
		ShortDescription: f.parseShortDescription(),
		LongDescription:  f.parseLongDescription(),
	}
}
