package main

import (
	"fmt"
	"time"

	"github.com/mugayoshi/vocabulary-helper/api/quotes"
	"github.com/mugayoshi/vocabulary-helper/config"
	"github.com/mugayoshi/vocabulary-helper/webhook"
)

type Cmd func()

func runCommand(f Cmd) {
	start := time.Now()
	f()
	fmt.Printf("execution time: %s\n", time.Since(start))
}

type QuoteInfo struct {
	quote      string
	originator string
}

func translateQuote(target string) QuoteInfo {
	rapidApiKey := config.GetEnvVariable("X_RAPID_API_KEY")
	quote, originator := quotes.GetQuote(rapidApiKey, target)
	return QuoteInfo{
		quote:      quote,
		originator: originator,
	}

}

func sendQuoteMessageToSlack() {
	info := translateQuote("es")
	url := config.GetEnvVariable("SLACK_WEBHOOK_LANG")
	message := fmt.Sprintf("Quote:\n%s\nby %s\n", info.quote, info.originator)
	webhook.SendMessageToMoneyChannel(url, message)
}

func main() {
	runCommand(sendQuoteMessageToSlack)
}
