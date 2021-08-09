package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Link struct {
	URL string `json:"URL"`
}

func GetLinks(site string) ([]Link, error) {
	links := make([]Link, 0)
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		link := Link{URL: href}
		links = append(links, link)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	if err := c.Visit(site); err != nil {
		return links, err
	}

	return links, nil
}
