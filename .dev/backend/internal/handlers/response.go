package handlers

import (
	"encoding/json"
	"net/http"
)

type StandardResponse struct {
	Success bool  	`json:"success"`
	Data interface{} `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Error string `json:"error,omitempty"`
}

func WriteJSONResponse(w http.ResponseWriter, status int, data interface{}) {
}