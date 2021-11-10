package server

import (
	"image"
	"image/png"
	"io/ioutil"
	"strconv"
	"tokamak/src/generator"
	miscgenerator "tokamak/src/generator/misc"
	profilegenerator "tokamak/src/generator/profile"
	"tokamak/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/nfnt/resize"
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

		switch p.Type {
		case "default":
			c.Set("Content-Type", "image/png")
			img = profilegenerator.RenderDefaultProfile(gen, p)
		case "modern":
			c.Set("Content-Type", "image/png")

			img = profilegenerator.RenderModernProfile(gen, p)
		case "profile_2":
			c.Set("Content-Type", "image/webp")

			img = profilegenerator.RenderProfileTwo(gen, p)
		default:
			c.Set("Content-Type", "image/png")
			img = profilegenerator.RenderDefaultProfile(gen, p)
		}

		if c.Query("w", "0") != "0" {
			if c.Query("h", "0") != "0" {
				parseW, err := strconv.ParseUint(c.Query("w", "0"), 10, 32)
				if err != nil {
					return err
				}
				parseH, err := strconv.ParseUint(c.Query("h", "0"), 10, 32)
				if err != nil {
					return err
				}

				if c.Query("type", "0") != "0" {
					switch c.Query("type", "nil") {
					case "thumb":
						newImage := resize.Thumbnail(uint(parseW), uint(parseH), img, resize.Lanczos3)
						img = newImage
					case "resize":
						newImage := resize.Resize(uint(parseW), uint(parseH), img, resize.Lanczos3)
						img = newImage
					default:
						c.Set("Content-Type", "text/plain; charset=utf-8")
						return c.SendString("😳 Oops! You forgot something your cute.")
					}
				}
			

			}
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
