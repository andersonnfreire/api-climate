package utils

import "strings"

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
