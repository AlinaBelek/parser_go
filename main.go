package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()
	c.OnHTML(".item-list", func(el *colly.HTMLElement) {
		el.ForEach("li", func(_ int, el *colly.HTMLElement) {

			title := el.ChildText("div.views-field-title > span.field-content")
			price := el.ChildText("div.views-field-display-price > span.field-content")

			writer.Write([]string{
				title,
				price,
			})
		})
		fmt.Println("Scrapping Complete")
	})
	c.Visit("https://kupislona-store.ru/catalog/aksessuary")
}
