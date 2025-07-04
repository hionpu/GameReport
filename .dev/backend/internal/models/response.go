package models

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty'`
	Message string      `json:"message,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func WriteSuccess(w http.ResponseWriter, data interface{}) error {
	return WriteJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    data,
	})
}

func WriteError(w http.ResponseWriter, status int, message string) error {
	return WriteJSON(w, status, APIResponse{
		Success: false,
		Message: message,
	})
}

func IsHTMXRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
