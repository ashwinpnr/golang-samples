package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Options struct {
	Password string
	LoginUrl string
}

type api struct {
	Options Options
	Client  http.Client
}

type PageRespose struct {
	Page string `json:"page"`
}

type WordsRespose struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func (w WordsRespose) GetResponse() string {
	return fmt.Sprintf("Words : %s\n", strings.Join(w.Words, ", "))
}

type OccurrenceRespose struct {
	Words map[string]int `json:"words"`
}

func (o OccurrenceRespose) GetResponse() string {
	words := []string{}
	for word, occurrence := range o.Words {
		words = append(words, fmt.Sprintf("%s : %d", word, occurrence))
	}
	return fmt.Sprintf("Occurrence : %s\n", strings.Join(words, ", "))
}

func (a api) DoGetRespose(requestURL string) (Response, error) {

	if _, err := url.ParseRequestURI(requestURL); err != nil {
		return nil, fmt.Errorf("url incorrect %s", requestURL)
	}

	response, errCode := a.Client.Get(requestURL)

	if errCode != nil {
		return nil, fmt.Errorf("Http Get Failed Error Code : %d", errCode)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("Response read failed : %s", err)
	}

	//fmt.Printf("Http Status Code : %d \n Respose : %s", response.StatusCode, string(body))
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid HTTP Respose Code : %d \n", response.StatusCode)
	}
	var pageRespose PageRespose
	err = json.Unmarshal(body, &pageRespose)
	if err != nil {
		//return nil, fmt.Errorf("Page response Marshalling failed  : %s \n", err)
		return nil, RequestError{
			Err:        fmt.Sprintf("Page unmarshal error: %s", err),
			StatusCode: response.StatusCode,
			Body:       string(body),
		}

	}

	switch pageRespose.Page {
	case "words":
		var wordsRespose WordsRespose

		err = json.Unmarshal(body, &wordsRespose)
		if err != nil {
			return nil, RequestError{
				Err:        fmt.Sprintf("Word unmarshal error: %s", err),
				StatusCode: response.StatusCode,
				Body:       string(body),
			}
		}

		return wordsRespose, nil
	case "occurrence":
		var occurrenceRespose OccurrenceRespose

		err = json.Unmarshal(body, &occurrenceRespose)
		if err != nil {
			return nil, RequestError{
				Err:        fmt.Sprintf("Occurrence unmarshal error: %s", err),
				StatusCode: response.StatusCode,
				Body:       string(body),
			}
		}

		return occurrenceRespose, nil

	default:
		return nil, fmt.Errorf("Invalid URL Respose Code : %s \n", requestURL)

	}

}
