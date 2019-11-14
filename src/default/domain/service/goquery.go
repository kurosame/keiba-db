package service

import (
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// GetHTMLDoc is get HTML document
func GetHTMLDoc(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("HTTP status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

// GetLinks is get all links in HTML document
func GetLinks(doc *goquery.Document, filterRegexp string) []string {
	var links []string
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		if regexp.MustCompile(filterRegexp).MatchString(link) {
			links = append(links, link)
		}
	})

	return links
}
