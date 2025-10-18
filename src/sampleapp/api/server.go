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

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.CreateEvent(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Printf("Starting API Server on port [%s]", port)
	return http.ListenAndServe(":"+port, nil)
}
