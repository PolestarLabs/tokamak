package server

import (
  "github.com/gofiber/fiber/v2"
  "tokamak/src/generator"
  "image"
  "image/png"
  "tokamak/src/generator/profile"
)

func StartServer (port string) {
  app := fiber.New()
  gen := generator.NewGenerator()
  encoder := png.Encoder {
    CompressionLevel: -1, // no compression.
  }
  
  app.Get("/version", func(c *fiber.Ctx) error {
    return c.SendString("tokamak v1 (fasthttp/fiber; golang)")
  })
  
  app.Post("/render/profile", func (c *fiber.Ctx) error {
    p := new(profilegenerator.ProfileData)
    
    if err := c.BodyParser(p); err != nil {
      return err
    }
    
    var img image.Image
    c.Set("Content-Type", "image/png")
    
    switch p.Type {
      case "default":
        img = profilegenerator.RenderDefaultProfile(gen, p)
      default:
        img = profilegenerator.RenderDefaultProfile(gen, p)
    }
    
    return encoder.Encode(c.Context(), img)
  })
  
  app.Listen(":" + port)
}