package words

import (
	"encoding/json"
	"log"
)

type randomWordRes struct {
	Word string `json:"word"`
}

// writes a random word to a channel
func GetRandomWord(apiKey string, isDebug bool, c chan<- string) {
	body := callAPiBase("?random=true", apiKey, isDebug)
	var result randomWordRes
	err := json.Unmarshal(body, &result)

	if err != nil {
		log.Panicln("unmarshal error")
	}
	c <- result.Word
}
