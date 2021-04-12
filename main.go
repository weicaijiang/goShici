package main

import (
	"goshici/gofish"
	"goshici/handle"
	"log"
)

func main() {
	authors := "https://so.gushiwen.org/authors"
	h := handle.AuthorHandle{}
	fish := gofish.NewGoFish()
	request, err := gofish.NewRequest("GET", authors, gofish.UserAgent, &h, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fish.Request = request
	fish.Visit()
}
