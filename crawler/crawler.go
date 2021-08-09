package crawler

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

type Link struct {
	URL string `json:"URL"`
}

func GetLinks(site string, domain string) ([]Link, error) {
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

	if domain != "" {
		links = filterByDomain(links, domain)
	}
	return links, nil
}

func filterByDomain(links []Link, domain string) []Link {
	r := make([]Link, 0)
	for _, v := range links {
		u, err := url.ParseRequestURI(v.URL)
		if err != nil {
			log.Println(err)
			continue
		}
		//log.Println(u.Host)
		if strings.Contains(u.Host, domain) {
			r = append(r, v)
		}
	}

	return r
}
