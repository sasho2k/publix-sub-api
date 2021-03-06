package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sasho2k/publix-sub-api/workforce"
	"net/http"
)

// StartService will serve as our router service. It is responsible for capturing traffic to our route and returning
// a response based on the parameter given to :storeNumber. There is also a check beforehand.
func StartService() {
	app := fiber.New(fiber.Config{GETOnly: true})

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	app.Get("/get-sale/", func(c *fiber.Ctx) error {
		sub, err := workforce.GrabDailySub(1295)
		if err != nil {
			return c.SendString(err.Error())
		}

		return c.JSON(sub)
	})

	app.Get("/get-subs/:storeNumber", func(c *fiber.Ctx) error {
		storeNumber, err := CheckParam(c.Params("storeNumber"))
		if err != nil {
			return c.SendString(err.Error())
		}

		subs, err := workforce.GrabSubs(storeNumber)
		if err != nil {
			return c.SendString(err.Error())
		}

		return c.JSON(subs)
	})

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
