package registry

import (
	"testing"

	"github.com/zohurul88/go-event/dispatcher"
	"github.com/zohurul88/go-event/event"
)

type UserCreatedEvent struct {
	event.BaseEvent
	Username string
}

func (e UserCreatedEvent) GetName() string {
	return e.EventName
}

func TestDispatcherRegistry_SetAndGetDispatcher(t *testing.T) {
	registry := NewDispatcherRegistry()

	userCreatedDispatcher := dispatcher.NewEventDispatcher[UserCreatedEvent]()
	registry.SetDispatcher("UserCreated", userCreatedDispatcher)

	retrieved := registry.GetDispatcher("UserCreated")
	if retrieved == nil {
		t.Error("expected to retrieve the dispatcher, got nil")
	}

	if retrieved != userCreatedDispatcher {
		t.Error("expected retrieved dispatcher to be the same as the set dispatcher")
	}
}
