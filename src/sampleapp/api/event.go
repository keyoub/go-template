package api

import (
	"encoding/json"
	"net/http"
	"sampleapp/service/capture"
)

type EventHandler struct {
	service capture.CaptureService
}

func NewEventHandler(service capture.CaptureService) *EventHandler {
	return &EventHandler{service: service}
}

type CreateEventRequest struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var req CreateEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CaptureEvent(req.Name, req.Data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "created"})
}
