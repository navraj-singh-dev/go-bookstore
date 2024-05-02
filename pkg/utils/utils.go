package util

import (
	"encoding/json"
	"io"
	"net/http"
)

// takes request's body and unmarshall it
func UnmarshallBody(r *http.Request, x interface{}) error {
	/*
		Flow:
		req body -> convert to []byte -> Unmarshall into given variable -> return nil
	*/

	// convert request's body to []byte
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// close the req body to save resources
	defer r.Body.Close()

	// un-marshall this []byte (body) in to the x variable
	if err := json.Unmarshal(body, x); err != nil {
		return err
	}

	return nil
}
