package fiber

import (
	"github.com/zohurul88/goevent"
)

func main() {
	InitApp()

	onReqDispatcher, _ := goevent.GetDispatcher[OnRequestEvent]("fiber.OnRequest")
	onReqDispatcher.Subscribe("fiber.OnRequest", func(e OnRequestEvent) {
		e.Ctx.SendString("Hello, World 👋!")
	}, 1)

}
