package handle

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
)

type AuthorHandle struct {
}

func (h *AuthorHandle) Worker(body io.Reader, url string) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
		return
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
