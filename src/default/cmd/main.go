package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	const baseURL = "https://db.netkeiba.com"
	const term = 3

	now := time.Now()

	for i := 0; i < term; i++ {
		target := now.AddDate(0, -i, 0)

		url := fmt.Sprintf("%s/?pid=race_top&date=%d%d", baseURL, target.Year(), target.Month())

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

		var links []string
		doc.Find("a").Each(func(_ int, s *goquery.Selection) {
			link, _ := s.Attr("href")

			if strings.Contains(link, "/race/list/") {
				links = append(links, link)
			}
		})

		for _, l := range links {
			time.Sleep(1 * time.Second)

			res, err := http.Get(baseURL + l)
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

			var links []string
			doc.Find("a").Each(func(_ int, s *goquery.Selection) {
				link, _ := s.Attr("href")

				if regexp.MustCompile(`^/race/\d+/$`).MatchString(link) {
					links = append(links, link)
				}
			})

			for _, l := range links {
				fmt.Println(baseURL + l)
			}
		}
	}
}
