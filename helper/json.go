package helper

import (
	"encoding/json"
	"net/http"
)

// Decode JSON request in a selected struct
func DecodeJsonRequest(r *http.Request, interfaceRef interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(interfaceRef); err != nil {
		return err
	}

	return nil
}
