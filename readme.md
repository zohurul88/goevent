# go-event

`goevent` is a flexible and easy-to-use event handling system for Go, designed to manage event subscriptions and dispatching with priority and asynchronous capabilities. This package provides a structured way to handle various events, making it suitable for building event-driven applications.

## Features

- **Event Interface**: Define your own events by implementing the `Event` interface.
- **Priority-Based Handlers**: Register event handlers with different priorities to control the order of execution.
- **Synchronous and Asynchronous Dispatch**: Dispatch events both synchronously and asynchronously.
- **Dispatcher Registry**: Manage multiple event dispatchers through a central registry, allowing for easy event management across your application.

## Installation

To use the `goevent` package in your Go project, run:

```sh
go get github.com/zohurul88/goevent
```

## Usage
Here's a basic example of how to use the `goevent` package:

```go
package order

import (
	"log"
	"net/http"
	"time"

	"github.com/zohurul88/goevent"
)

type OrderEvent struct {
	goevent.BaseEvent
	OrderID       string
	Amount        float64
	Items         []string
	CustomerID    string
	OrderDate     string
	LoyaltyPoints int
	Message       string
}

func (e OrderEvent) GetName() string {
	return e.EventName
}

func main() {
	reg := goevent.GetGlobalDispatcherRegistry()

	// Create and set dispatcher for OrderEvent
	orderDispatcher := goevent.NewEventDispatcher[OrderEvent]()
	reg.SetDispatcher("OrderEvent", orderDispatcher)

	// Subscribe handlers to the OrderEvent with different priorities
	orderDispatcher.Subscribe("OrderEvent", func(e OrderEvent) {
		log.Printf("Order confirmed: %s\n", e.OrderID)
	}, 1)

	orderDispatcher.Subscribe("OrderEvent", func(e OrderEvent) {
		log.Printf("Payment processed for order: %s, amount: %.2f\n", e.OrderID, e.Amount)
	}, 2)

	orderDispatcher.Subscribe("OrderEvent", func(e OrderEvent) {
		log.Printf("Inventory updated for order: %s, items: %v\n", e.OrderID, e.Items)
	}, 3)

	orderDispatcher.Subscribe("OrderEvent", func(e OrderEvent) {
		orderDate, _ := time.Parse("2006-01-02", e.OrderDate)
		shippingDate := orderDate.AddDate(0, 0, 3)
		log.Printf("Shipping scheduled for order: %s on date: %s\n", e.OrderID, shippingDate.String())
	}, 4)

	orderDispatcher.Subscribe("OrderEvent", func(e OrderEvent) {
		log.Printf("Customer notified for order: %s, message: %s\n", e.OrderID, e.Message)
	}, 5)

	orderDispatcher.Subscribe("OrderEvent", func(e OrderEvent) {
		log.Printf("Invoice generated for order: %s, invoice: %s\n", e.OrderID, "INV123")
	}, 6)

	orderDispatcher.Subscribe("OrderEvent", func(e OrderEvent) {
		log.Printf("Loyalty points updated for customer: %s, points: %d\n", e.CustomerID, e.LoyaltyPoints)
	}, 7)

	// Example HTTP server to handle order confirmation
	http.HandleFunc("/confirm", func(w http.ResponseWriter, r *http.Request) {
		orderID := r.URL.Query().Get("orderID")
		if orderID == "" {
			http.Error(w, "Order ID is required", http.StatusBadRequest)
			return
		}

		orderEvent := OrderEvent{
			BaseEvent:     goevent.BaseEvent{EventName: "OrderEvent"},
			OrderID:       orderID,
			Amount:        100.00,
			Items:         []string{"item1", "item2"},
			CustomerID:    "customer123",
			OrderDate:     "2024-06-01",
			LoyaltyPoints: 10,
			Message:       "Your order has been confirmed",
		}
		orderDispatcher.DispatchSync(orderEvent)
		// orderDispatcher.DispatchAsync(orderEvent, &reg.AsyncWG)

		w.Write([]byte("Order confirmed and events triggered"))
	})

	// Start the HTTP server
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}


```

## Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License.
