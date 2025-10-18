package api

import (
	"errors"
	"log"
	"net/http"
	"sampleapp/repo/event"
	"sampleapp/service/capture"
)

func StartServer(port string) error {
	if port == "" {
		return errors.New("port is required")
	}

	repo := event.NewMockEventRepo()
	service := capture.NewCaptureService(repo)
	handler := NewEventHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /events", handler.CreateEvent)

	log.Printf("Starting API Server on port [%s]", port)
	return http.ListenAndServe(":"+port, mux)
}
