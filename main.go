package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	// Define & parse flags
	search := flag.String("search", "", "search for quotes on a topic")
	author := flag.String("author", "", "search for quotes by author")
	tag := flag.String("tag", "", "search for quotes with tag")
	showTags := flag.Bool("showtags", false, "display tags for each quote")
	max := flag.Int("n", 10, "maximum number of results")

	flag.Parse()

	// Search or get QOTD depending on flags
	if *search != "" || *author != "" || *tag != "" {
		quotes, err := GetQuotes(*search, *author, *tag, *max)
		if err != nil {
			log.Fatal(err)
		}
		for i, quote := range quotes {
			printQuote(quote, *showTags)
			if i != len(quotes)-1 {
				fmt.Println()
			}
		}
	} else {
		quote, err := GetRandomQuote()
		if err != nil {
			log.Fatal(err)
		}
		printQuote(quote, *showTags)
	}
}
