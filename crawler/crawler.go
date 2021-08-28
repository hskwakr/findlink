package crawler

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Link struct {
	URL string `json:"url"`
}

func GetLinks(site string, domain string, o io.Writer) ([]Link, error) {
	links := make([]Link, 0)
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		link := Link{URL: href}
		links = append(links, link)
	})

	c.OnRequest(func(r *colly.Request) {
		msg := fmt.Sprint("Visiting: ", r.URL.String(), "\n")

		if _, err := o.Write([]byte(msg)); err != nil {
			log.Fatal("failed to write on request:", err)
		}
	})

	if err := c.Visit(site); err != nil {
		return links, fmt.Errorf("filed to visit: %w", err)
	}

	if domain != "" {
		links = filterByString(links, domain)
	}

	return links, nil
}

func filterByString(links []Link, str string) []Link {
	r := make([]Link, 0)

	for _, v := range links {
		if strings.Contains(v.URL, str) {
			r = append(r, v)
		}
	}

	return r
}
