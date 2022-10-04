package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"url-function/pkg/api"
)

// {"page":"words","input":"test3","words":["test1","test2","test3"]}

func main() {

	var (
		parseURL *url.URL
		err      error
		inputURL string
		password string
	)

	flag.StringVar(&inputURL, "url", "", "Input URL to access")
	flag.StringVar(&password, "password", "", "password to access")
	flag.Parse()

	if parseURL, err = url.ParseRequestURI(inputURL); err != nil {
		fmt.Printf("URL Validation Failed %s", err)
		flag.Usage()
		os.Exit(1)
	}

	loginURL := parseURL.Scheme + "://" + parseURL.Host + "/login"

	apiInstance := api.New(api.Options{
		Password: password,
		LoginUrl: loginURL,
	})

	response, err := apiInstance.DoGetRespose(parseURL.String())

	if err != nil {
		fmt.Printf("Error : %s", err)
		os.Exit(1)
	}
	if response == nil {
		fmt.Printf("No response : %s", err)
		os.Exit(1)
	}

	fmt.Printf("Response : %s", response.GetResponse())

}
