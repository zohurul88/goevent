package dispatcher

import (
	"sync"
	"testing"

	"github.com/zohurul88/go-event/event"
)

type UserCreatedEvent struct {
	event.BaseEvent
	Username string
}

func (e UserCreatedEvent) GetName() string {
	return e.EventName
}

func TestEventDispatcher_SubscribeAndDispatchSync(t *testing.T) {
	dispatcher := NewEventDispatcher[UserCreatedEvent]()
	var result []string
	dispatcher.Subscribe("UserCreated", func(event UserCreatedEvent) {
		result = append(result, "Handler1: "+event.Username)
	}, 1)
	dispatcher.Subscribe("UserCreated", func(event UserCreatedEvent) {
		result = append(result, "Handler2: "+event.Username)
	}, 2)

	ev := UserCreatedEvent{BaseEvent: event.BaseEvent{EventName: "UserCreated"}, Username: "johndoe"}
	dispatcher.DispatchSync(ev)

	expected := []string{"Handler2: johndoe", "Handler1: johndoe"}
	for i, r := range result {
		if r != expected[i] {
			t.Errorf("expected '%s', got '%s'", expected[i], r)
		}
	}
}

func TestEventDispatcher_DispatchAsync(t *testing.T) {
	dispatcher := NewEventDispatcher[UserCreatedEvent]()
	var mu sync.Mutex
	var result []string
	dispatcher.Subscribe("UserCreated", func(event UserCreatedEvent) {
		mu.Lock()
		result = append(result, "Handler1: "+event.Username)
		mu.Unlock()
	}, 1)
	dispatcher.Subscribe("UserCreated", func(event UserCreatedEvent) {
		mu.Lock()
		result = append(result, "Handler2: "+event.Username)
		mu.Unlock()
	}, 2)

	ev := UserCreatedEvent{BaseEvent: event.BaseEvent{EventName: "UserCreated"}, Username: "johndoe"}

	var wg sync.WaitGroup
	dispatcher.DispatchAsync(ev, &wg)
	wg.Wait()

	expected := []string{"Handler2: johndoe", "Handler1: johndoe"}
	for i, r := range result {
		if r != expected[i] {
			t.Errorf("expected '%s', got '%s'", expected[i], r)
		}
	}
}
