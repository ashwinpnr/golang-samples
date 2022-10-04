package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LoginRequest struct {
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func doLoginRequest(client http.Client, loginURL, password string) (string, error) {

	loginRequest := LoginRequest{
		Password: password,
	}
	body, err := json.Marshal(loginRequest)
	if err != nil {
		//return nil, fmt.Errorf("Page response Marshalling failed  : %s \n", err)
		return "", fmt.Errorf("login marshall error %s", err)
	}

	resp, err := client.Post(loginURL, "application/json", bytes.NewBuffer(body))

	if err != nil {
		//return nil, fmt.Errorf("Page response Marshalling failed  : %s \n", err)
		return "", RequestError{
			Err:        fmt.Sprintf("Login Request Error: %s", err),
			StatusCode: resp.StatusCode,
			Body:       "",
		}

	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("Response read failed : %s", err)
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Invalid HTTP Respose Code : %d \n", resp.StatusCode)
	}
	var loginResponse LoginResponse
	err = json.Unmarshal(body, &loginResponse)
	if err != nil {
		//return nil, fmt.Errorf("Page response Marshalling failed  : %s \n", err)
		return "", RequestError{
			Err:        fmt.Sprintf("Login Response unmarshal error: %s", err),
			StatusCode: resp.StatusCode,
			Body:       string(body),
		}

	}

	return loginResponse.Token, nil
}
