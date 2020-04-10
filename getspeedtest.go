package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	client := http.Client{}
	url := "https://speedtestserver.com/"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("request err", err)
	}
	rsp, err := client.Do(req)
	if err != nil {
		fmt.Println("rsp err")
	}
	defer rsp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		country := s.Find("td").Eq(0).Text()
		city := s.Find("td").Eq(1).Text()
		provider := s.Find("td").Eq(2).Text()
		host := s.Find("td").Eq(3).Text()
		fmt.Printf("Review %d: \n\t%s\n\t%s\n\t%s\n\t%s\n", i, country, city, provider, host)
	})
}
