package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
)

// URL https://quotes.toscrape.com/
func main() {
	// Create a new Colly collector
	collector := colly.NewCollector()

	// Define the URL you want to scrape
	url := "https://quotes.toscrape.com/page/1/"

	// Set up callbacks to handle scraping events
	collector.OnHTML(".quote", func(element *colly.HTMLElement) {
		//Extract data from HTML elements
		quote := element.ChildText("span.text")

		quote = strings.TrimSpace(quote)
		fmt.Println(quote)
	})

	err := collector.Visit(url)
	if err != nil {
		log.Fatal(err)
		return
	}

}
