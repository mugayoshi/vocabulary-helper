package main

import (
	"fmt"
	"time"

	"github.com/mugayoshi/vocabulary-helper/api/quotes"
	"github.com/mugayoshi/vocabulary-helper/api/translate"
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
	quote       string
	originator  string
	translation string
}

func translateQuote() QuoteInfo {
	rapidApiKey := config.GetEnvVariable("X_RAPID_API_KEY")
	quote, originator := quotes.GetQuote(rapidApiKey, "en")
	translation := translate.Translate(rapidApiKey, quote, "es", "en")
	fmt.Println(translation)
	return QuoteInfo{
		quote:       quote,
		originator:  originator,
		translation: translation,
	}

}

func sendQuoteMessageToSlack() {
	info := translateQuote()
	url := config.GetEnvVariable("SLACK_WEBHOOK_LANG")
	message := fmt.Sprintf("Quote: %s  by %s, \nTranslation: %s\n", info.quote, info.originator, info.translation)
	webhook.SendMessageToMoneyChannel(url, message)
}

func main() {
	runCommand(sendQuoteMessageToSlack)
}
