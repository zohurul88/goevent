package dispatcher

import (
	"sort"
	"sync"

	"github.com/zohurul88/go-event/event"
)

// EventDispatcher manages event subscriptions and dispatching with priorities.
type EventDispatcher[T event.Event] struct {
	handlers map[string][]event.PriorityHandler[T]
	mu       sync.RWMutex
}

// NewEventDispatcher creates a new EventDispatcher.
func NewEventDispatcher[T event.Event]() *EventDispatcher[T] {
	return &EventDispatcher[T]{
		handlers: make(map[string][]event.PriorityHandler[T]),
	}
}

// Subscribe registers a handler with a priority for an event.
func (d *EventDispatcher[T]) Subscribe(eventName string, handler event.EventHandler[T], priority int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers[eventName] = append(d.handlers[eventName], event.PriorityHandler[T]{Handler: handler, Priority: priority})
}

// DispatchSync dispatches an event to all registered handlers, sorted by priority, synchronously.
func (d *EventDispatcher[T]) DispatchSync(e T) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	handlers := d.handlers[e.GetName()]
	// Sort handlers by priority
	sort.SliceStable(handlers, func(i, j int) bool {
		return handlers[i].Priority > handlers[j].Priority
	})
	for _, ph := range handlers {
		ph.Handler(e)
	}
}

// DispatchAsync dispatches an event to all registered handlers, sorted by priority, asynchronously.
func (d *EventDispatcher[T]) DispatchAsync(e T, wg *sync.WaitGroup) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	handlers := d.handlers[e.GetName()]
	// Sort handlers by priority
	sort.SliceStable(handlers, func(i, j int) bool {
		return handlers[i].Priority > handlers[j].Priority
	})

	for _, ph := range handlers {
		wg.Add(1)
		go func(handler event.EventHandler[T], event T) {
			defer wg.Done()
			handler(event)
		}(ph.Handler, e)
	}
}
