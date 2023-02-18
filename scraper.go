package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName)

	if err != nil {
		log.Fatalf("error creating file %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()

	c.OnHTML(".prd", func(h *colly.HTMLElement) {
		writer.Write([]string{
			h.ChildText("h3.name"),
			// h.ChildText(""),
		})

		fmt.Println("Scrapping Complete")
	})

	for i := 0; i < 10; i++ {

		c.Visit("https://www.jumia.com.ng/flash-sales/?page=" + strconv.Itoa(i) + "#catalog-listing")
	}

	log.Println("Scraping complete")
	log.Println(c)
}
