package scrape

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/gocolly/colly"
)

type user struct {
	name string
	days []string
}

func Scrape(url string, cookie string) {
	c := colly.NewCollector()
	c.OnHTML(".privboard-row", printElement)

	c.Cookies(cookie)
	if err := c.Visit(url); err != nil {
		log.Error("cant visit scrape url")
	}
}

func printElement(e *colly.HTMLElement) {
	s, _ := e.DOM.Html()
	fmt.Printf("%s\n", s)

}

func getMemberlist() {

}
