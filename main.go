package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var destination string
var to_date, return_date string

func main() {
	// to_date = "2023-03-30"
	// return_date = "2023-04-06"
	// destination = "CPT-HKT"

	// url := "https://www.cheapflights.co.za/flight-search/CPT-HKT/2023-03-30/2023-04-06?sort=bestflight_a"
	// resp, err := http.Get(url)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer resp.Body.Close()

	// // Parse the HTML response body using goquery
	// doc, err := goquery.NewDocumentFromReader(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // Find all elements with class "f8F1-price-text"
	// prices := doc.Find(".f8F1")
	// fmt.Println(prices)
	// // Iterate over each element and print its text

	// prices.Each(func(i int, s *goquery.Selection) {
	// 	fmt.Println(s.Text())
	// })
	// Define the URL to scrape
	url := "https://www.cheapflights.co.za/flight-search/CPT-HKT/2023-03-30/2023-04-06?sort=bestflight_a"

	// Send an HTTP GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// Parse the HTML response body using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Extract the best flight prices from the webpage
	prices := []string{}
	doc.Find(".f8F1-price-text").Each(func(i int, s *goquery.Selection) {
		// Clean up the text by removing any whitespace characters
		price := strings.TrimSpace(s.Text())
		if price != "" {
			prices = append(prices, price)
		}
	})

	// Print out the best flight prices
	if len(prices) > 0 {
		fmt.Println("Best flight prices:")
		for _, price := range prices {
			fmt.Println(price)
		}
	} else {
		fmt.Println("No flight prices found")
	}

}
