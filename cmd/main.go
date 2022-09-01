package main

import (
	"fmt"

	"github.com/mugayoshi/vocabulary-helper/api/words"
	"github.com/mugayoshi/vocabulary-helper/config"
)

func main() {
	rapidApiKey := config.GetEnvVariable("X_RAPID_API_KEY")
	randomWord := words.GetRandomWord(rapidApiKey, false)
	fmt.Printf("random word is %s\n", randomWord)
}
