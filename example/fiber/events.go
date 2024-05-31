package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zohurul88/goevent"
)

type OnRequestEvent struct {
	goevent.BaseEvent
	Ctx *fiber.Ctx
}

type OnResponseEvent struct {
	goevent.BaseEvent
	Ctx *fiber.Ctx
}

func (e OnRequestEvent) GetName() string {
	return e.EventName
}

func (e OnResponseEvent) GetName() string {
	return e.EventName
}

// modify wet a new function
func NewOnRequestEvent() OnRequestEvent {
	return OnRequestEvent{
		BaseEvent: goevent.BaseEvent{EventName: "fiber.OnRequest"},
	}
}

func NewOnResponseEvent() OnResponseEvent {
	return OnResponseEvent{
		BaseEvent: goevent.BaseEvent{EventName: "fiber.OnRequest"},
	}
}
