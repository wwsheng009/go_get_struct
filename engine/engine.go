package engine

import "strings"

func GetCamelCase(input string) string {
	arr := strings.Split(input, "_")
	var result string
	for _, v := range arr {
		result += strings.Title(v)
	}
	return result
}
