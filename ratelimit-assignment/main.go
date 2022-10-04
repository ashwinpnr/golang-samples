package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func hitUrl(urlString string, wait *sync.WaitGroup) {
	response, err := http.Get(urlString)
	var returnString string
	if err != nil {
		returnString = fmt.Sprintf("Http Get Error : %v", err)
	}

	if response.StatusCode != 200 {
		returnString = fmt.Sprintf("Invalid Response %v", response.StatusCode)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		returnString = fmt.Sprintf("response body parse error Error : %v", err)
	}

	returnString = fmt.Sprintf("Body: %s\n", string(body))
	fmt.Printf("%s\n", returnString)
	wait.Done()
}

func main() {

	var wait sync.WaitGroup
	rateLimitUrl := "http://localhost:8080/ratelimit"
	//c := make(chan string)

	for j := 1; j <= 10; j++ {
		time.Sleep(999 * time.Millisecond)
		for i := 1; i <= 5; i++ {
			wait.Add(1)
			go hitUrl(rateLimitUrl, &wait)

		}
	}
	wait.Wait()

}
