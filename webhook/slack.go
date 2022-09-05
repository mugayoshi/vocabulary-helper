package webhook

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func SendMessageToMoneyChannel(url, message string) {
	messageValue := fmt.Sprintf(`{"text":"%s"}`, message)
	var jsonStr = []byte(messageValue)
	req, newReqErr := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if newReqErr != nil {
		log.Panicln("NewRequest error", newReqErr)
	}

	client := &http.Client{}
	resp, doErr := client.Do(req)
	if doErr != nil {
		log.Panicln("Do error", doErr)
	}
	defer resp.Body.Close()
	log.Println("response Status:", resp.Status)
	body, readAllErr := ioutil.ReadAll(resp.Body)
	if readAllErr != nil {
		log.Panicln("ReadAll error", readAllErr)
	}
	log.Println("response Body:", string(body))
	if strings.Contains(resp.Status, "200") {
		log.Println("successfully sent a message to Slack.")
	}
}
