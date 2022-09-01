package words

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// call Word API
// path shouldn't contain a slash at the beginning.
func callAPiBase(path, key string, isDebug bool) []byte {
	base := "https://wordsapiv1.p.rapidapi.com/words"
	url := strings.Join([]string{base, path}, "/")
	if isDebug {
		log.Printf("URL: %s\n", url)
	}
	req, newReqErr := http.NewRequest("GET", url, nil)

	if newReqErr != nil {
		log.Panicln("NewRequest error")
	}

	req.Header.Add("X-RapidAPI-Key", key)
	req.Header.Add("X-RapidAPI-Host", "wordsapiv1.p.rapidapi.com")

	res, doErr := http.DefaultClient.Do(req)
	if doErr != nil {
		log.Panicln("HTTP Client.Do Error")
	}

	defer res.Body.Close()
	body, readAllErr := ioutil.ReadAll(res.Body)

	if readAllErr != nil {
		log.Panicln("ReadAll err")
	}
	if isDebug {
		fmt.Println(res)
		fmt.Println(string(body))
	}
	return body
}
