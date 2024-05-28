package event

// Event is the interface for all events.
type Event interface {
	GetName() string
}

type BaseEvent struct {
	EventName string
}
