package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response code:", r.StatusCode)
	})

	c.OnHTML("footer", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit("https://gowebexamples.com/")
}
