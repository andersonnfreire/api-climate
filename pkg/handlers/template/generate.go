package template

import (
	"errors"

	"github.com/andersonnfreire/api-climate/pkg/handlers/prevision"
)

// Gera um HTML com base nos dados da previsão do tempo.
func RenderTemplateHTML(weatherData *prevision.WeatherForecastsResponse) (string, error) {
	if len(weatherData.ClimaticWeather) == 0 {
		return "", errors.New("nenhuma informação meteorológica foi encontrada para esta cidade")
	}

	// urlImagem := fmt.Sprintf("https://openweathermap.org/img/wn/%s@2x.png", weatherData.ClimaticWeather[0].Icon)
	return "", nil
}
