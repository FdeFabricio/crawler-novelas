package crawler

import (
	"fmt"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/fdefabricio/crawler-novelas/model"
	"github.com/fdefabricio/crawler-novelas/validator"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

var (
	infoC, actorsC *colly.Collector
	mutex          = &sync.Mutex{}
)

func init() {
	infoC = colly.NewCollector(
		colly.AllowedDomains("pt.wikipedia.org"),
		colly.CacheDir("./crawlercache"),
		//colly.Async(true),
	)

	actorsC = infoC.Clone()
}

func Run(urls []string) (novelas map[string]*model.Novela) {
	novelas = make(map[string]*model.Novela)
	warningCount := 0

	// On every <a> element in a cell representing the title of a novela, fill in the info and call the actors collector
	infoC.OnHTML("table:has(tr th:contains('Título'))", func(e *colly.HTMLElement) {
		e.ForEach("tr:has(td i a)", func(i int, e *colly.HTMLElement) {
			name := e.ChildText("td:nth-child(4) i a")
			url := e.Request.AbsoluteURL(e.ChildAttr("td:nth-child(4) i a", "href"))
			info := model.BasicInfo{
				Authors:   strings.Split(e.ChildText("td:nth-child(6)"), "\n"),
				Chapters:  e.ChildText("td:nth-child(5)"),
				Directors: strings.Split(e.ChildText("td:nth-child(7)"), "\n"),
				Hour:      strings.Split(fmt.Sprintf("%s", e.Request.URL), "#")[1],
				Name:      name,
				Year:      e.ChildText("td:nth-child(2)")[len(e.ChildText("td:nth-child(2)"))-4:],
				URL:       url,
			}

			// TODO: hardcoded because Wikipedia's entry is misspelled
			if name == "Brega e Chique" {
				name = "Brega & Chique"
			}

			mutex.Lock()
			novelas[strings.ToLower(name)] = &model.Novela{info, make([]string, 0)}
			mutex.Unlock()

			actorsC.Visit(url)
		})
	})

	// On every page of a novela, extract the name of the actors - table(s) after <h2>Elenco</h2>
	actorsC.OnHTML("body", func(e *colly.HTMLElement) {
		name := e.ChildText("h1 i")
		actors := make([]string, 0)
		e.ForEach("h2:contains('Elenco')", func(_ int, e *colly.HTMLElement) {
			e.DOM.NextFilteredUntil("table", "h2").Each(func(i int, s *goquery.Selection) {
				s.Find("table:not(:has(table)) tbody").Each(func(j int, s *goquery.Selection) {
					s.Find("tr td:first-child").Each(func(i int, s *goquery.Selection) {
						actors = append(actors, s.Text())
					})
				})
			})
		})

		if novelas[strings.ToLower(name)] == nil {
			log.Errorf("not found: %s", name)
			return
		}

		mutex.Lock()
		novelas[strings.ToLower(name)].AppendActors(actors)
		mutex.Unlock()

		errs := validator.Check(*novelas[strings.ToLower(name)])
		for _, err := range errs {
			if err != nil {
				log.Warnf("%s: %s", name, err)
				warningCount += 1
			}
		}
	})

	infoC.OnRequest(func(r *colly.Request) {
		log.Debug("[INFORM]", r.URL.String())
	})

	actorsC.OnRequest(func(r *colly.Request) {
		log.Debug("[ACTORS]", r.URL.String())
	})

	for _, url := range urls {
		err := infoC.Visit(url)
		if err != nil {
			log.Error(err)
		}
	}

	//infoC.Wait()
	//actorsC.Wait()

	fmt.Printf("%d validation warnings found\n", warningCount)

	return
}
