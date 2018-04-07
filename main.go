package main

import (
	"fmt"
	"os"

	"github.com/fdefabricio/crawler-novelas/crawler"
	log "github.com/sirupsen/logrus"
)

func init() {
	switch os.Getenv("ENV") {
	case "VALIDATION":
		log.SetLevel(log.WarnLevel)
		f, err := os.Create("validation.log")
		if err != nil {
			log.Fatal("Couldn't create/open validation log file")
		}
		log.SetOutput(f)
	default:
		log.SetLevel(log.ErrorLevel)
	}
}

func main() {
	listOfListOfNovelas := []string{
		"https://pt.wikipedia.org/wiki/Lista_de_telenovelas_das_seis_da_Rede_Globo#6",
		"https://pt.wikipedia.org/wiki/Lista_de_telenovelas_das_sete_da_Rede_Globo#7",
		"https://pt.wikipedia.org/wiki/Lista_de_telenovelas_das_oito_da_Rede_Globo#8",
		"https://pt.wikipedia.org/wiki/Lista_de_telenovelas_das_nove_da_Rede_Globo#9",
		"https://pt.wikipedia.org/wiki/Lista_de_telenovelas_das_dez_da_Rede_Globo#10",
		"https://pt.wikipedia.org/wiki/Lista_de_telenovelas_das_onze_da_Rede_Globo#11",
	}

	novelas := crawler.Run(listOfListOfNovelas)

	fmt.Printf("%d novelas scrapped\n", len(novelas))
}
