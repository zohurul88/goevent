package goevent

import (
	"fmt"
)

type UserCreatedEvent struct {
	BaseEvent
	Username string
}

func (e UserCreatedEvent) GetName() string {
	return e.EventName
}

type UserDeletedEvent struct {
	BaseEvent
	Username string
}

func (e UserDeletedEvent) GetName() string {
	return e.EventName
}

func main() {
	registry := GetGlobalDispatcherRegistry()

	// Create and set dispatcher for UserCreatedEvent
	userCreatedDispatcher := NewEventDispatcher[UserCreatedEvent]()
	registry.SetDispatcher("UserCreated", userCreatedDispatcher)

	// Subscribe handlers to the UserCreated event with different priorities
	userCreatedDispatcher.Subscribe("UserCreated", func(event UserCreatedEvent) {
		fmt.Printf("User created: %s\n", event.Username)
	}, 1)

	userCreatedDispatcher.Subscribe("UserCreated", func(event UserCreatedEvent) {
		fmt.Printf("Sending welcome email to: %s\n", event.Username)
	}, 2)

	// Create and set dispatcher for UserDeletedEvent
	userDeletedDispatcher := NewEventDispatcher[UserDeletedEvent]()
	registry.SetDispatcher("UserDeleted", userDeletedDispatcher)

	// Subscribe handlers to the UserDeleted event with different priorities
	userDeletedDispatcher.Subscribe("UserDeleted", func(event UserDeletedEvent) {
		fmt.Printf("User deleted: %s\n", event.Username)
	}, 1)

	userDeletedDispatcher.Subscribe("UserDeleted", func(event UserDeletedEvent) {
		fmt.Printf("Sending account deletion email to: %s\n", event.Username)
	}, 2)

	// Dispatch a UserCreated event synchronously
	userCreatedEvent := UserCreatedEvent{
		BaseEvent: BaseEvent{EventName: "UserCreated"},
		Username:  "johndoe",
	}
	userCreatedDispatcher.DispatchSync(userCreatedEvent)

	// Dispatch a UserDeleted event asynchronously
	userDeletedEvent := UserDeletedEvent{
		BaseEvent: BaseEvent{EventName: "UserDeleted"},
		Username:  "johndoe",
	}
	userDeletedDispatcher.DispatchAsync(userDeletedEvent, &registry.AsyncWG)

	// Wait for all async events to complete
	registry.WaitForAsyncCompletion()
}
