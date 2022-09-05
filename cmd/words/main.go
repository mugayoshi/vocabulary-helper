package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/mugayoshi/vocabulary-helper/api/words"
	"github.com/mugayoshi/vocabulary-helper/config"
)

func getRandomWords(apiKey string) []string {
	randomWords := make([]string, 5)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		c := make(chan string)
		wg.Add(1)
		go func(c chan<- string) {
			defer wg.Done()
			words.GetRandomWord(apiKey, false, c)
		}(c)

		randomWords[i] = <-c
		close(c)
	}
	wg.Wait()

	return randomWords
}
func main() {
	rapidApiKey := config.GetEnvVariable("X_RAPID_API_KEY")
	randomWords := getRandomWords(rapidApiKey)
	fmt.Printf("%s\n", strings.Join(randomWords, ","))
}
