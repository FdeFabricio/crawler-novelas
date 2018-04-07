package main

import (
	"github.com/fdefabricio/crawler-novelas/crawler"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.ErrorLevel)
	log.SetLevel(log.WarnLevel)
	log.SetLevel(log.InfoLevel)
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

	log.Infof("%d novelas scrapped", len(novelas))
}
