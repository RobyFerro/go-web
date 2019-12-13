package helper

import (
	"encoding/json"
	"net/http"
)

func DecodeJsonRequest(r *http.Request, interfaceRef interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(interfaceRef); err != nil {
		return err
	}

	return nil
}
