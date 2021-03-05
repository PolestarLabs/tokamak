package server

import (
	"github.com/gofiber/fiber/v2"
	"image"
	"image/png"
	"io/ioutil"
	"tokamak/src/generator"
	"tokamak/src/generator/misc"
	"tokamak/src/generator/profile"
	"tokamak/src/utils"
)

func StartServer(port string) {
	app := fiber.New()
	gen := generator.NewGenerator()
	encoder := png.Encoder{
		CompressionLevel: -1, // no compression.
	}

	app.Static("/static", "../assets/images")
	app.Get("/version", func(c *fiber.Ctx) error {
		return c.SendString("tokamak v1.1-dev1 (fasthttp/fiber; golang)")
	})
	app.Get("/get_asset_list/:image", func(c *fiber.Ctx) error {
		if _, found := utils.Find([]string{"foundation", "badges", "stickers", "bgs"}, c.Params("image")); !found {
			return c.SendString("no 💋")
		}

		files, err := ioutil.ReadDir("../assets/images/" + c.Params("image"))
		if err != nil {
			panic(err)
			return c.SendString("[]")
		}

		return c.JSON(utils.FilterFileList(files))
	})

	app.Post("/render/profile", func(c *fiber.Ctx) error {
		p := new(profilegenerator.ProfileData)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		var img image.Image
		c.Set("Content-Type", "image/png")

		switch p.Type {
			case "default":
				img = profilegenerator.RenderDefaultProfile(gen, p)
			case "modern":
				img = profilegenerator.RenderModernProfile(gen, p)
				break;
			default:
				img = profilegenerator.RenderDefaultProfile(gen, p)
		}

		return encoder.Encode(c.Context(), img)
	})




	app.Post("/render/license", func(c *fiber.Ctx) error {
		p := new(miscgenerator.LicenseData)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		c.Set("Content-Type", "image/png")

		return encoder.Encode(c.Context(), miscgenerator.RenderLicenseImage(gen, p))
	})

	app.Post("/render/rize", func(c *fiber.Ctx) error {
		p := new(miscgenerator.RizeData)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		c.Set("Content-Type", "image/png")

		return encoder.Encode(c.Context(), miscgenerator.RenderRizeImage(gen, p))
	})

	app.Post("/render/laranjo", func(c *fiber.Ctx) error {
		p := new(miscgenerator.LaranjoData)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		c.Set("Content-Type", "image/png")

		return encoder.Encode(c.Context(), miscgenerator.RenderLaranjoImage(gen, p))
	})
	app.Listen(":" + port)
}
