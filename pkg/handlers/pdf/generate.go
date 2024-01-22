package pdf

import (
	"errors"
	"fmt"

	"github.com/andersonnfreire/api-climate/pkg/handlers/prevision"
	"github.com/andersonnfreire/api-climate/pkg/utils"
	"github.com/jung-kurt/gofpdf"
)

// GeneratePDF gera um PDF com base nos dados da previsão do tempo.
func GeneratePDFHandler(weatherData *prevision.WeatherForecastsResponse) (*gofpdf.Fpdf, error) {

	if len(weatherData.ClimaticWeather) == 0 {
		return nil, errors.New("nenhuma informação meteorológica foi encontrada para esta cidade")
	}

	urlImagem := fmt.Sprintf("https://openweathermap.org/img/wn/%s@2x.png", weatherData.ClimaticWeather[0].Icon)
	img, err := utils.DownloadImage(urlImagem)
	if err != nil {
		return nil, fmt.Errorf("não foi possivel baixar a imagem remetente ao tempo dessa cidade:%s", err.Error())
	}

	// Salvar a imagem temporariamente
	tempImagePath := "temp_image.png"
	err = utils.SaveImage(img, tempImagePath)
	if err != nil {
		return nil, fmt.Errorf("erro ao salvar a imagem temporária:%s", err.Error())
	}

	// Criar um novo PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Adicionar uma nova página
	pdf.AddPage()

	// Definir a fonte
	pdf.SetFont("Arial", "", 12)

	// Adicionar título
	addTitle(pdf, utils.Utf8ToIso("PREVISÃO DO TEMPO"))

	// Coordenadas iniciais para a imagem
	x, _, _, _ := pdf.GetMargins()

	// Adicionar imagem
	addImage(pdf, tempImagePath, x, 15, 40, 40)

	cityValues := ConvertValuesResponseAPIWeather(weatherData)

	// Adicionar a tabela ao PDF
	for _, label := range cityLabels {
		addRow(pdf, utils.Utf8ToIso(label), GetCityValueByLabel(cityValues, label))
	}

	if err := utils.RemoveImage(tempImagePath); err != nil {
		return nil, err
	}

	return pdf, nil
}

// addTitle adiciona um título ao PDF com espaçamento após o título
func addTitle(pdf *gofpdf.Fpdf, title string) {
	pdf.Cell(0, 10, title)
	pdf.Ln(40)
}

// addImage adiciona uma imagem ao PDF
func addImage(pdf *gofpdf.Fpdf, imagePath string, x, y, width, height float64) {
	pdf.ImageOptions(imagePath, x, y, width, height, false, gofpdf.ImageOptions{ImageType: "PNG"}, 0, "")
}

// addRow adiciona uma linha de dados à tabela
func addRow(pdf *gofpdf.Fpdf, values ...string) {
	columnWidth := 40.0
	for i, value := range values {
		// Usar CellFormat para centralizar somente os valores, não os rótulos
		align := "L" // Por padrão, alinhar à esquerda para rótulos
		if i%2 == 1 {
			align = "C" // Centralizar os valores (índice ímpar)
		}
		pdf.CellFormat(columnWidth, 10, value, "1", 0, align, false, 0, "")
	}
	pdf.Ln(10)
}
