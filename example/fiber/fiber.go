// add go fiber example here
package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zohurul88/goevent"
)

func InitApp() {
	onRequestEvent := NewOnRequestEvent()
	onResponseEvent := NewOnResponseEvent()
	onRequestDispatcher := goevent.NewEventDispatcher[OnRequestEvent]()
	onResponseDispatcher := goevent.NewEventDispatcher[OnResponseEvent]()
	goevent.GetGlobalDispatcherRegistry().SetDispatcher(onRequestEvent.GetName(), onRequestDispatcher)
	goevent.GetGlobalDispatcherRegistry().SetDispatcher(onResponseEvent.GetName(), onResponseDispatcher)

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		onRequestEvent.Ctx = c
		onRequestDispatcher.DispatchSync(onRequestEvent)
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Use(func(c *fiber.Ctx) error {
		onResponseEvent.Ctx = c
		onResponseDispatcher.DispatchSync(onResponseEvent)
		return c.Next()
	})
	app.Listen(":3000")
}
