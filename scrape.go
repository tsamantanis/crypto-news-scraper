package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
)

type Data struct {
	Heading string
	Link    string
}

var boldRed *color.Color = color.New(color.FgRed, color.Bold)

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every h4 element call callback
	var arr []*Data
	c.OnHTML("h4", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(_ int, kf *colly.HTMLElement) {
			dataTemp := &Data{
				Heading: e.Text,
				Link:    "https://cryptonews.com" + e.ChildAttr("a", "href"),
			}
			arr = append(arr, dataTemp)
			temp, tempErr := json.MarshalIndent(dataTemp, "", "  ")
			handleError((tempErr))
			fmt.Print(string(temp))
			fmt.Println()
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://cryptonews.com/
	c.Visit("https://cryptonews.com/")

	js, jsonErr := json.MarshalIndent(arr, "", "\t")
	handleError(jsonErr)
	f, fileErr := os.OpenFile("output.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0660)
	handleError(fileErr)

	defer f.Close()
	if !json.Valid([]byte(js)) {
		boldRed.Print("Error: ")
		fmt.Print("The JSON data produced by the scraper is invalid")
	}
	f.WriteString(string(js))
}

func handleError(err error) {
	if err != nil {
		boldRed.Print("Error:")
		fmt.Print(err)
		panic(err)
	}
}
