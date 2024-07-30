package utils

import (
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
	"strings"
)

// Função para converter uma string para maiúsculas
func ToUpper(str string) string {
	return strings.ToUpper(str)
}

func Utf8ToIso(s string) string {
	var iso8859_1 []byte
	for _, runeValue := range s {
		if runeValue <= 0xFF {
			iso8859_1 = append(iso8859_1, byte(runeValue))
		}
	}
	return string(iso8859_1)
}

func DownloadImage(url string) (image.Image, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func SaveImage(img image.Image, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}

func RemoveImage(tempImagePath string) error {
	if err := os.Remove(tempImagePath); err != nil {
		return fmt.Errorf("erro ao remover a imagem:%s", err.Error())
	}
	return nil
}
