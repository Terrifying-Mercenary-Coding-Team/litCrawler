package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type result struct {
	url    string
	status int
}

func hitURL(url string, channel chan<- result) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		channel <- result{url: url, status: 0}
	} else {
		channel <- result{url: url, status: 1}
	}
}

func main() {
	// Check for valid input
	if len(os.Args) != 2 {
		panic("Error: Invalid input")
	}

	// Decode json array
	urls := []string{}
	err := json.Unmarshal([]byte(os.Args[1]), &urls)
	if err != nil {
		panic("Error: Decoding Json")
	}

	// Process hitURL
	channel := make(chan result)
	results := map[string]int{}
	for _, url := range urls {
		go hitURL(url, channel)
	}

	// Wait for receiving
	for i := 0; i < len(urls); i++ {
		tmp := <-channel
		results[tmp.url] = tmp.status
	}
	// Encode result
	jsonResults, err := json.Marshal(results)
	if err != nil {
		panic("Error: Encoding Json")
	}
	fmt.Println(string(jsonResults))
}
