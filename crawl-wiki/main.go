package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"

)

func main() {
	res, err := http.Get("https://vi.wikipedia.org/wiki/Trang_Ch%C3%ADnh")
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
	title:= doc.Find("#main-on-this-day h2.main-box__title a").Text()
	fmt.Printf("%s\n", title)
	
	doc.Find("#main-on-this-day .main-box__body ul li").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		item := s.Text()
		fmt.Printf("%d: %s\n", i, item)
	})
	
}
