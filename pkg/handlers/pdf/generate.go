package pdf

import (
	"github.com/andersonnfreire/api-climate/pkg/handlers/prevision"
	"github.com/jung-kurt/gofpdf"
)

// GeneratePDF gera um PDF com base nos dados da previsão do tempo.
func GeneratePDF(weatherData *prevision.WeatherForecastsResponse) (*gofpdf.Fpdf, error) {
	// Criar um novo PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Adicionar uma nova página
	pdf.AddPage()

	// Definir a fonte
	pdf.SetFont("Arial", "", 12)

	// Adicionar título
	pdf.Cell(0, 10, "PREVISÃO DO TEMPO")
	pdf.Ln(35) // Adicionar espaçamento após o título

	// Coordenadas iniciais para a imagem
	_, y, _, _ := pdf.GetMargins()

	// Obter largura e altura da página
	pageWidth, _ := pdf.GetPageSize()

	// Calcular a largura da imagem
	imageWidth := 40.0 // Substitua pela largura real da sua imagem

	// Calcular a posição x para centralizar a imagem
	_ = (pageWidth - float64(imageWidth)) / 2

	// Adicionar imagem
	imagePath := "imagem.png" // Substitua pelo caminho da sua imagem
	imageOptions := gofpdf.ImageOptions{ImageType: "PNG"}
	pdf.ImageOptions(imagePath, 30, y, imageWidth, 0, false, imageOptions, 0, "")

	// Dados da cidade
	data := [][]string{
		{"Cidade", "Vitoria"},
		{"Latitude", "-10"},
		{"Longitude", "-710"},
		{"Céu", "Nublado"},
		{"Precipitação", "Chuvisco"},
		{"Vento", "Fraco a Moderado"},
		{"Temp: Máx", "31º"},
		{"Temp: Min", "26º"},
	}

	// Adicionar a tabela ao PDF
	for _, row := range data {
		AddRow(pdf, row...)
	}

	return pdf, nil
}

// Adicionar uma linha de dados à tabela
func AddRow(pdf *gofpdf.Fpdf, values ...string) {
	for _, value := range values {
		pdf.CellFormat(40, 10, value, "1", 0, "L", false, 0, "")
	}
	pdf.Ln(10)
}
