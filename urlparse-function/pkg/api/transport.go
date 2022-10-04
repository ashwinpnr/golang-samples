package api

import (
	"fmt"
	"net/http"
	"os"
)

type MyJWTClient struct {
	token     string
	transport http.RoundTripper
	password  string
	loginUrl  string
}

func (m *MyJWTClient) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.token == "" {
		if m.password != "" {

			token, err := doLoginRequest(http.Client{}, m.loginUrl, m.password)
			if err != nil {
				fmt.Printf("Login Error  : %s", err)
				os.Exit(1)
			}
			if token == "" {
				fmt.Printf("Token received is blank")
				os.Exit(1)
			}
			fmt.Printf("Token received : %s", token)
			m.token = token
		}

	}

	if m.token != "" {

		req.Header.Add("Authorization", "Bearer "+m.token)
	}
	return m.transport.RoundTrip(req)
}
