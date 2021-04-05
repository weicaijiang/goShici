package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	authors := "https://so.gushiwen.org/authors"
	res, err := http.Get(authors)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("error status:", res.Status, " code:", res.StatusCode)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".sons").Find(".cont").Find("a").Each(func(i int, s *goquery.Selection) {
		author := s.Text()
		link, exists := s.Attr("href")
		if !exists {
			log.Println("not exit author:", author)
			return
		}
		log.Println("author:", author, " link:", link)
	})
}
