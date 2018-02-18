package main

import "github.com/bobesa/chalk"

func printQuote(quote Quote, showTags bool) {
	if showTags {
		tagString := chalk.White().Faint().Sprint(quote.Tags)
		chalk.Cyan().Italic().Printf("%s %s\n", quote.Body, tagString)
	} else {
		chalk.Cyan().Italic().Printf("%s\n", quote.Body)
	}
	if quote.Author != "" {
		chalk.Magenta().Printf("%s\n", quote.Author)
	}
}
