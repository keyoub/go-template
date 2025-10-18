package capture

import (
	"sampleapp/entity"
	"sampleapp/repo/event"
)

type CaptureService interface {
	CaptureEvent(name, data string) error
}

type captureService struct {
	repo event.EventRepo
}

func NewCaptureService(repo event.EventRepo) CaptureService {
	return &captureService{repo: repo}
}

func (s *captureService) CaptureEvent(name, data string) error {
	event := &entity.Event{Name: name, Data: data}
	return s.repo.Create(event)
}
