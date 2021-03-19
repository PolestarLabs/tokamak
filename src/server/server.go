package server

import (
	"github.com/gofiber/fiber/v2"
	"image/png"
	"io/ioutil"
	"tokamak/src/dynamicgen"
)

func StartServer(port string) {
	app := fiber.New()
	encoder := png.Encoder{
		CompressionLevel: -1, // no compression.
	}
	dg := dynamicgen.New()

	app.Static("/static", "../assets/images")

	app.Get("/version", func(c *fiber.Ctx) error {
		return c.SendString("tokamak v2 (fasthttp/fiber; golang)")
	})

	app.Get("/get_asset_list/:image", func(c *fiber.Ctx) error {
		if _, found := Find([]string{"foundation", "badges", "stickers", "bgs"}, c.Params("image")); !found {
			return c.SendString("no 💋")
		}

		files, err := ioutil.ReadDir("../assets/images/" + c.Params("image"))
		if err != nil {
			return c.SendString("[]")
		}

		return c.JSON(FilterFileList(files))
	})

	app.Post("/render/:image", func(c *fiber.Ctx) error {
		if c.Params("image") == "" {
			return c.SendString("{\"error\":true}")
		}

		p := make(map[string]interface{})
		if err := c.BodyParser(&p); err != nil {
			return err
		}

		c.Set("Content-Type", "image/png")

		img := dg.Render(c.Params("image"), p)
		if img == nil {
			c.Set("Content-Type", "application/json")
			return c.SendString("{\"error\":true}")
		}

		return encoder.Encode(c.Context(), img)
	})

	app.Get("/reload_blueprints", func(c *fiber.Ctx) error {
		prints := dynamicgen.LoadBlueprints()
		dg.Blueprints = prints

		return c.SendString("ok")
	})

	app.Listen(":" + port)
}
