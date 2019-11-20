package service

import (
	"bufio"
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
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

	bodyToJp := transform.NewReader(bufio.NewReader(res.Body), japanese.EUCJP.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(bodyToJp)
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

// GetTexts is get all texts matched elements by selector
func GetTexts(doc *goquery.Document, selector string) []string {
	var texts []string
	doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
		texts = append(texts, s.Text())
	})

	return texts
}
