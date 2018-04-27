package utils

import (
	"regexp"
	"strings"
)

func IsIn(slice []string, b string) bool {
	for _, a := range slice {
		if a == b {
			return true
		}
	}
	return false
}

func ConvertSpecialCaracteres(s string) string {
	return strings.Replace(s, "%26", "&", -1)
}

func ExtractNameElencoDeURL(s string) string {
	return strings.TrimSpace(strings.Split(strings.Replace(strings.Split(s, "Elenco_de_")[1], "_", " ", -1), "(")[0])
}

// "O Astro (2011)" --> "O Astro"
// "143 [nota 8]" --> "143"
func PruneName(s string) string {
	return strings.TrimSpace(regexp.MustCompile(`\[.*?\]|\(.*?\)`).ReplaceAllString(s, ""))
}
