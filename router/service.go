package router

import (
	"github.com/gofiber/fiber"
	"github.com/sasho2k/publix-sub-api/workforce"
)

// StartService will serve as our router service. It is responsible for capturing traffic to our route and returning
// a response based on the parameter given to :storeNumber. There is also a check beforehand.
func StartService() {
	app := fiber.New(fiber.Config{StrictRouting: true, GETOnly: true})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	app.Get("/get-sale/:storeNumber", func(c *fiber.Ctx) error {
		storeNumber, err := CheckParam(c.Params("storeNumber"))
		if err != nil {
			return err
		}

		sub, err := workforce.GrabDailySub(storeNumber)
		if err != nil {
			return err
		}

		return c.JSON(sub)
	})

	app.Get("/get-subs/:storeNumber", func(c *fiber.Ctx) error {
		storeNumber, err := CheckParam(c.Params("storeNumber"))
		if err != nil {
			return err
		}

		subs, err := workforce.GrabSubs(storeNumber)
		if err != nil {
			return err
		}

		return c.JSON(subs)
	})

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
