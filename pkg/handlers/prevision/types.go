package prevision

type WeatherForecastsResponse struct {
	Coordinates         Coordinates `json:"coord"`
	ClimaticWeather     []Weather   `json:"weather"`
	InfoMain            InfoMain    `json:"main"`
	InfoWind            InfoWind    `json:"wind"`
	InfoClouds          InfoClouds  `json:"clouds"`
	TimeDataCalculation int         `json:"dt"`
	Sys                 InfoCountry `json:"sys"`
	Name                string      `json:"name"`
	Timezone            int         `json:"timezone"`
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
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type InfoWind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type InfoClouds struct {
	All int `json:"all"`
}

type InfoCountry struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}
