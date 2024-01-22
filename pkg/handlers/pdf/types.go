package pdf

import (
	"fmt"

	"github.com/andersonnfreire/api-climate/pkg/handlers/prevision"
	"github.com/andersonnfreire/api-climate/pkg/utils"
)

type CityValues struct {
	City              string
	Latitude          string
	Longitude         string
	ConditionClimatic ConditionClimatic
	Pressure          string
	Humidity          string
	Cloudiness        string
	Temp              string
	TempMax           string
	TempMin           string
}

type ConditionClimatic struct {
	Name        string
	Description string
}

var cityLabels = []string{
	"Cidade",
	"Latitude",
	"Longitude",
	"Céu",
	"Precipitação",
	"Pressão",
	"Umidade",
	"Nebulosidade",
	"Temperatura",
	"Temp: Máx",
	"Temp: Min",
}

// CityValuesMap é um mapa que mapeia rótulos para funções que obtêm os valores correspondentes de CityValues.
var CityValuesMap = map[string]func(*CityValues) string{
	"Cidade":       func(c *CityValues) string { return utils.Utf8ToIso(c.City) },
	"Latitude":     func(c *CityValues) string { return c.Latitude },
	"Longitude":    func(c *CityValues) string { return c.Longitude },
	"Céu":          func(c *CityValues) string { return utils.Utf8ToIso(c.ConditionClimatic.Name) },
	"Precipitação": func(c *CityValues) string { return utils.Utf8ToIso(c.ConditionClimatic.Description) },
	"Pressão":      func(c *CityValues) string { return c.Pressure },
	"Umidade":      func(c *CityValues) string { return c.Humidity },
	"Nebulosidade": func(c *CityValues) string { return c.Cloudiness },
	"Temperatura":  func(c *CityValues) string { return c.Temp },
	"Temp: Máx":    func(c *CityValues) string { return c.TempMax },
	"Temp: Min":    func(c *CityValues) string { return c.TempMin },
}

func ConvertValuesResponseAPIWeather(weatherData *prevision.WeatherForecastsResponse) *CityValues {
	return &CityValues{
		City:      weatherData.Name,
		Latitude:  fmt.Sprintf("%.2f", weatherData.Coordinates.Lat),
		Longitude: fmt.Sprintf("%.2f", weatherData.Coordinates.Lon),
		ConditionClimatic: ConditionClimatic{
			Name:        weatherData.ClimaticWeather[0].Main,
			Description: weatherData.ClimaticWeather[0].Description,
		},
		Pressure:   fmt.Sprintf("%d", weatherData.InfoMain.Pressure),
		Humidity:   fmt.Sprintf("%d", weatherData.InfoMain.Humidity),
		Cloudiness: fmt.Sprintf("%d", weatherData.InfoClouds.All),
		Temp:       fmt.Sprintf("%.2f", weatherData.InfoMain.Temp),
		TempMax:    fmt.Sprintf("%.2f", weatherData.InfoMain.TempMax),
		TempMin:    fmt.Sprintf("%.2f", weatherData.InfoMain.TempMin),
	}
}

// GetCityValueByLabel obtém o valor correspondente ao rótulo do CityValues usando o mapa.
func GetCityValueByLabel(cityValues *CityValues, label string) string {
	if getValueFunc, exists := CityValuesMap[label]; exists {
		return getValueFunc(cityValues)
	}
	return "Não informado"
}
