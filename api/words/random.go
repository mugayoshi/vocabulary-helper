package words

import (
	"encoding/json"
	"log"
)

type randomWordResResult struct {
	Definition string `json:"definition"`
}

type randomWordRes struct {
	Word    string                `json:"word"`
	Results []randomWordResResult `json:"results"`
}

func GetRandomWord(apiKey string, isDebug bool) string {
	body := callAPiBase("?random=true", apiKey, isDebug)
	var result randomWordRes
	err := json.Unmarshal(body, &result)

	if err != nil {
		log.Panicln("unmarshal error")
	}
	return result.Word
}
