package fiber

import (
	"github.com/zohurul88/goevent"
)

func main() {

	onReqDispatcher, _ := goevent.GetDispatcher[OnRequestEvent]("fiber.OnRequest")
	onReqDispatcher.Subscribe("fiber.OnRequest", func(e OnRequestEvent) {
		e.Ctx.SendString("Hello, World 👋!")
	}, 1)

	InitApp()

}
