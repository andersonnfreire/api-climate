package prevision

type ParamsGetPrevision struct {
	Token    string `validate:"required"`
	Cidade   string `validate:"required"`
	Language string
}

type WeatherForecastsResponse struct {
	Coordinates     Coordinates `json:"coord"`
	ClimaticWeather []Weather   `json:"weather"`
	InfoMain        InfoMain    `json:"main"`
	InfoClouds      InfoClouds  `json:"clouds"`
	Name            string      `json:"name"`
}

type Coordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type InfoMain struct {
	Temp     float64 `json:"temp"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
	Pressure int64   `json:"pressure"`
	Humidity int64   `json:"humidity"`
}

type InfoClouds struct {
	All int64 `json:"all"`
}
