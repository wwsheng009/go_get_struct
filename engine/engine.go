package engine

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetCamelCase(input string) string {
	input1 := strings.ReplaceAll(input, "-", "_")

	arr := strings.Split(input1, "_")
	var result string
	caser := cases.Title(language.AmericanEnglish)
	for _, v := range arr {
		result += caser.String(v)
	}

	return result
}
