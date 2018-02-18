package main

import (
	"fmt"
	"github.com/dghubble/sling"
	"os"
)

var (
	quoteApi      = sling.New().Base("https://favqs.com/api/")
	qotdPath      = "qotd.json"
	quotesPath    = "quotes"
	defaultApiKey = "1501a5726d3be7d5839dc9085ba47848"
	favqsApiKey   string
)

func init() {
	var ok bool
	favqsApiKey, ok = os.LookupEnv("FAVQS_API_KEY")
	if !ok {
		favqsApiKey = defaultApiKey
	}
}

type Quote struct {
	Id     int      `json:"id"`
	Body   string   `json:"body"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
}

type QOTDResponse struct {
	Quote Quote `json:"quote"`
}

type QuotesResponse struct {
	Quotes []Quote `json:"quotes"`
}

// Quotes params
type QuoteParams struct {
	Filter string `url:"filter,omitempty"`
	Type   string `url:"type,omitempty"`
	Page   int    `url:"page,omitempty"`
}

func GetRandomQuote() (Quote, error) {
	qotdResponse := QOTDResponse{}
	_, err := quoteApi.Get(qotdPath).ReceiveSuccess(&qotdResponse)
	return qotdResponse.Quote, err
}

func GetQuotes(search, author, tag string, max int) ([]Quote, error) {
	quotesResponse := QuotesResponse{}
	quotesQuery := QuoteParams{
		Filter: search,
	}
	if author != "" {
		quotesQuery.Filter = author
		quotesQuery.Type = "author"
	}
	if tag != "" {
		quotesQuery.Filter = tag
		quotesQuery.Type = "tag"
	}
	_, err := quoteApi.Get(quotesPath).Add("Authorization", fmt.Sprintf("Token token=%s", favqsApiKey)).QueryStruct(&quotesQuery).ReceiveSuccess(&quotesResponse)
	if err != nil {
		return nil, err
	}

	quotes := quotesResponse.Quotes
	if max < len(quotes) {
		quotes = quotes[:max]
	}
	return quotes, nil
}
