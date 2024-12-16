package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload any, w http.ResponseWriter) error {

	if r.Body == nil {
		w.Write([]byte("The body is nill please enter the body parameters"))
		return fmt.Errorf("The body is empty cant be empty")
	}
	return json.NewDecoder(r.Body).Decode(&payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
