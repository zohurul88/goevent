package event

import (
	"testing"
)

// UserCreatedEvent represents a user creation event.
type UserCreatedEvent struct {
	BaseEvent
	Username string
}

// GetName returns the name of the event.
func (e UserCreatedEvent) GetName() string {
	return e.EventName
}

// UserDeletedEvent represents a user deletion event.
type UserDeletedEvent struct {
	BaseEvent
	Username string
}

// GetName returns the name of the event.
func (e UserDeletedEvent) GetName() string {
	return e.EventName
}

func TestUserCreatedEvent_GetName(t *testing.T) {
	event := UserCreatedEvent{BaseEvent: BaseEvent{EventName: "UserCreated"}, Username: "johndoe"}
	if event.GetName() != "UserCreated" {
		t.Errorf("expected event name to be 'UserCreated', got '%s'", event.GetName())
	}
}

func TestUserDeletedEvent_GetName(t *testing.T) {
	event := UserDeletedEvent{BaseEvent: BaseEvent{EventName: "UserDeleted"}, Username: "johndoe"}
	if event.GetName() != "UserDeleted" {
		t.Errorf("expected event name to be 'UserDeleted', got '%s'", event.GetName())
	}
}
