package openweathermap

type OpenWeatherMapEndpoint struct {
	Data []DataProperty `json:"data"`
}

type DataProperty struct {
	Temperature float32           `json:"temp"`
	FeelsLike   float32           `json:"feels_like"`
	Humidity    float32           `json:"humidity"`
	Visibility  float32           `json:"visibility"`
	WindSpeed   float32           `json:"wind_speed"`
	Weather     []WeatherProperty `json:"weather"`
}

type WeatherProperty struct {
	ShortDescription string `json:"main"`
	LongDescription  string `json:"description"`
}
