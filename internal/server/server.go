package server

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/vladimish/client-bot/internal/server/handlers"
	"github.com/vladimish/client-bot/internal/server/requests"
)

func StartApi() {
	app := fiber.New()

	app.Post("/add_table", func(c *fiber.Ctx) error {
		req := requests.AddTable{}
		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			return err
		}

		return handlers.HandleAddTable(req)
	})

	app.Post("/delete_table", func(c *fiber.Ctx) error {
		req := requests.DeleteTable{}
		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			return err
		}

		return handlers.HandleDeleteTable(req)
	})

	app.Post("/get_tables", func(c *fiber.Ctx) error {
		res, err := handlers.HandleGetTables()
		if err != nil {
			return err
		}
		bytes, err := json.Marshal(res)
		if err != nil {
			return err
		}
		_, err = c.Write(bytes)
		if err != nil {
			return err
		}

		return nil
	})

	err := app.Listen(":1721")
	if err != nil {
		panic(err)
	}
}
