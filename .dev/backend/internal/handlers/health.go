package handlers

import (
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":  "ok",
		"service": "insightGG-platform",
	}

	WriteJSONResponse(w, http.StatusOK, response)
}
