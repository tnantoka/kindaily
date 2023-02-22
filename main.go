package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func scrape(host string, path string) {
	res, err := http.Get(host + path)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#ebooks-deals-storefront-0 .a-carousel-card").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Find("a").Attr("href")
		if !exists {
			return
		}

		title, exists := s.Find("img").Attr("alt")
		if exists {
			fmt.Print("\033[1m\033[4m")
			fmt.Printf("%d: %s", i + 1, title)
			fmt.Println("\033[0m")
			fmt.Printf("%s%s\n\n", host, href)
		}
	})
}

func main() {
	scrape("https://www.amazon.co.jp", "/b?node=3338926051")
}
