# go-event

`go-event` is a flexible and easy-to-use event handling system for Go, designed to manage event subscriptions and dispatching with priority and asynchronous capabilities. This package provides a structured way to handle various events, making it suitable for building event-driven applications.

## Features

- **Event Interface**: Define your own events by implementing the `Event` interface.
- **Priority-Based Handlers**: Register event handlers with different priorities to control the order of execution.
- **Synchronous and Asynchronous Dispatch**: Dispatch events both synchronously and asynchronously.
- **Dispatcher Registry**: Manage multiple event dispatchers through a central registry, allowing for easy event management across your application.

## Installation

To use the `go-event` package in your Go project, run:

```sh
go get github.com/zohurul88/go-event
```

## Usage
Here's a basic example of how to use the `go-event` package:

```go
package main

import (
	"fmt"

	"github.com/zohurul88/go-event/dispatcher"
	"github.com/zohurul88/go-event/event"
	"github.com/zohurul88/go-event/registry"
)

type UserCreatedEvent struct {
	event.BaseEvent
	Username string
}

func (e UserCreatedEvent) GetName() string {
	return e.EventName
}

type UserDeletedEvent struct {
	event.BaseEvent
	Username string
}

func (e UserDeletedEvent) GetName() string {
	return e.EventName
}

func main() {
	registry := registry.GetGlobalDispatcherRegistry()

	// Create and set dispatcher for UserCreatedEvent
	userCreatedDispatcher := dispatcher.NewEventDispatcher[UserCreatedEvent]()
	registry.SetDispatcher("UserCreated", userCreatedDispatcher)

	// Subscribe handlers to the UserCreated event with different priorities
	userCreatedDispatcher.Subscribe("UserCreated", func(event UserCreatedEvent) {
		fmt.Printf("User created: %s\n", event.Username)
	}, 1)

	userCreatedDispatcher.Subscribe("UserCreated", func(event UserCreatedEvent) {
		fmt.Printf("Sending welcome email to: %s\n", event.Username)
	}, 2)

	// Create and set dispatcher for UserDeletedEvent
	userDeletedDispatcher := dispatcher.NewEventDispatcher[UserDeletedEvent]()
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
		BaseEvent: event.BaseEvent{EventName: "UserCreated"},
		Username:  "johndoe",
	}
	userCreatedDispatcher.DispatchSync(userCreatedEvent)

	// Dispatch a UserDeleted event asynchronously
	userDeletedEvent := UserDeletedEvent{
		BaseEvent: event.BaseEvent{EventName: "UserDeleted"},
		Username:  "johndoe",
	}
	userDeletedDispatcher.DispatchAsync(userDeletedEvent, &registry.AsyncWG)

	// Wait for all async events to complete
	registry.WaitForAsyncCompletion()
}

```

## Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License.
