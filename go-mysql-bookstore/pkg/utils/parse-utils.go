package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseBody(r *http.Request, input any) error {
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return fmt.Errorf("error parsing : %s", err)
	}
	return nil
}
