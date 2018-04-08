package utils

import "strings"

func IsIn(slice []string, b string) bool {
	for _, a := range slice {
		if a == b {
			return true
		}
	}
	return false
}

// "O Astro (2011)" --> "O Astro"
// "143 [nota 8]" --> "143"
func TrimSuffix(s string) string {
	return strings.TrimSpace(strings.Split(strings.Split(strings.Split(s, "_(")[0], "(")[0], "[")[0])
}

func ConvertSpecialCaracteres(s string) string {
	return strings.Replace(s, "%26", "&", -1)
}

func ExtractNameElencoDeURL(s string) string {
	return strings.TrimSpace(strings.Split(strings.Replace(strings.Split(s, "Elenco_de_")[1], "_", " ", -1), "(")[0])
}
