package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)

type Data struct {
	Heading string
	Link    string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every h4 element call callback
	c.OnHTML("h4", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(_ int, kf *colly.HTMLElement) {
			dataTemp := Data{
				Heading: e.Text,
				Link:    "https://cryptonews.com" + e.ChildAttr("a", "href"),
			}
			js, err := json.MarshalIndent(dataTemp, "", "    ")
			if err != nil {
				fmt.Print(err)
				panic(err)
			}
			fmt.Print(string(js))
			// fmt.Println(js.Link)
			fmt.Println()
		})

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://cryptonews.com/
	c.Visit("https://cryptonews.com/")
}
