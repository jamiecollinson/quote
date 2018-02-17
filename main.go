package main

import (
	"flag"
	"github.com/bobesa/chalk"
	"github.com/dghubble/sling"
	"log"
)

const quotesUrl string = "https://favqs.com/api/"
const qotd string = "qotd.json"

type Quote struct {
	Id     int      `json:"id"`
	Body   string   `json:"body"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
}

type QOTDResponse struct {
	Quote Quote `json:"quote"`
}

func main() {
	// qotd := flag.Bool("qotd", false, "display quote of the day")

	flag.Parse()

	quotesBase := sling.New().Base(quotesUrl)

	qotdResponse := QOTDResponse{}
	_, err := quotesBase.Get(qotd).ReceiveSuccess(&qotdResponse)
	if err != nil {
		log.Fatal(err)
	}

	chalk.Cyan().Print(qotdResponse.Quote.Body)
	chalk.Magenta().Italic().Print(" ", qotdResponse.Quote.Author, "\n")
}
