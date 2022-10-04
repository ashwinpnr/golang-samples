package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// {"page":"words","input":"test3","words":["test1","test2","test3"]}

type PageRespose struct {
	Page string `json:"page"`
}

type WordsRespose struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}

type OccurrenceRespose struct {
	Words map[string]int `json:"words"`
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage incorrect : ./url_parse <url>")
		os.Exit(1)
	}

	inputURL := args[1]

	if _, err := url.ParseRequestURI(inputURL); err != nil {
		fmt.Printf("url incorrect %s", inputURL)
		os.Exit(1)
	}

	response, errCode := http.Get(inputURL)

	if errCode != nil {
		log.Fatal(errCode)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("Http Status Code : %d \n Respose : %s", response.StatusCode, string(body))
	if response.StatusCode != 200 {
		fmt.Printf("Invalid HTTP Respose Code : %d \n", response.StatusCode)
		os.Exit(1)
	}
	var pageRespose PageRespose
	err = json.Unmarshal(body, &pageRespose)
	if err != nil {
		log.Fatal(err)
	}

	switch pageRespose.Page {
	case "words":
		var wordsRespose WordsRespose

		err = json.Unmarshal(body, &wordsRespose)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("words : %v", wordsRespose.Words)
	case "occurrence":
		var occurrenceRespose OccurrenceRespose

		err = json.Unmarshal(body, &occurrenceRespose)
		if err != nil {
			log.Fatal(err)
		}

		for key, value := range occurrenceRespose.Words {
			fmt.Printf("Word : %s , Occurrence : %d \n", key, value)
		}

	default:
		fmt.Printf("Invalid url")
	}

}
