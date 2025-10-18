package event

import (
	"sampleapp/entity"
	"time"

	"github.com/google/uuid"
)

type EventRepo interface {
	Create(event *entity.Event) error
}

type mockEventRepo struct {
	events map[string]*entity.Event
}

func NewMockEventRepo() EventRepo {
	return &mockEventRepo{events: make(map[string]*entity.Event)}
}

func (r *mockEventRepo) Create(event *entity.Event) error {
	if event.ID == "" {
		event.ID = uuid.New().String()
	}
	event.CreatedAt = time.Now()
	r.events[event.ID] = event
	return nil
}
