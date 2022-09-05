package translate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type TranslatedText struct {
	Text string `json:"translatedText"`
}

type TranslatesInRes struct {
	Translations []TranslatedText `json:"translations"`
}

type GoogleTranslateResponse struct {
	Data TranslatesInRes `json:"data"`
}

func callGoogleTranslateApi(apiKey, input, target, source string) []byte {
	url := "https://google-translate1.p.rapidapi.com/language/translate/v2"

	query := func(input string) string {
		q := strings.ReplaceAll(input, " ", "%20")
		return fmt.Sprintf("q=%s&target=%s&source=%s", q, target, source)
	}(input)
	payload := strings.NewReader(query)

	req, newReqErr := http.NewRequest("POST", url, payload)
	if newReqErr != nil {
		log.Panicln("NewRequest error")
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, doErr := http.DefaultClient.Do(req)
	if doErr != nil {
		log.Panicln("Do error")
	}

	defer res.Body.Close()
	body, readAllErr := ioutil.ReadAll(res.Body)
	if readAllErr != nil {
		log.Panicln("ReadAll error")
	}

	return body
}

func Translate(apiKey, input, target, source string) string {
	body := callGoogleTranslateApi(apiKey, input, target, source)
	var result GoogleTranslateResponse
	err := json.Unmarshal(body, &result)

	if err != nil {
		log.Panicln("unmarshal error")
	}
	if len(result.Data.Translations) == 0 {
		return ""
	}

	return result.Data.Translations[0].Text
}
