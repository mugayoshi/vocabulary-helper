package quotes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type originator struct {
	Name string `json:"name"`
}

type QuoteResponse struct {
	Language   string     `json:"language_code"`
	Content    string     `json:"content"`
	Originator originator `json:"originator"`
}

func callQuotesApi(apiKey, lang string) []byte {
	url := fmt.Sprintf("https://quotes15.p.rapidapi.com/quotes/random/?language_code=%s", lang)

	req, newReqErr := http.NewRequest("GET", url, nil)
	if newReqErr != nil {
		log.Panicln("NewRequest error", newReqErr)
	}

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "quotes15.p.rapidapi.com")

	res, doErr := http.DefaultClient.Do(req)
	if doErr != nil {
		log.Panicln("Do error", doErr)
	}
	defer res.Body.Close()
	body, readAllErr := ioutil.ReadAll(res.Body)
	if readAllErr != nil {
		log.Panicln("ReadAll error", readAllErr)
	}
	return body
}

// returns a random quote and its originator
func GetQuote(apiKey, lang string) (string, string) {
	var result QuoteResponse
	body := callQuotesApi(apiKey, lang)
	err := json.Unmarshal(body, &result)
	if err != nil {
		log.Panicln("unmarshal error")
	}
	return result.Content, result.Originator.Name
}
