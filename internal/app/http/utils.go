package http

import (
	"encoding/json"
	"errors"
	"net/http"
)

func GetJSON(target *interface{}, res *http.Response) error {
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return errors.New("internal server error")
	}

	return nil
}
