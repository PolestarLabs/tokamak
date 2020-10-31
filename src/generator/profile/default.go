package profilegenerator

import (
  "image"
  "tokamak/src/generator"
)

import "github.com/fogleman/gg"

type DefaultProfile struct {
  Generator generator.Generator
  AvatarURL, BackgroundURL string
  Name, AboutMe, FavColor, HighestRole string
}

func (p DefaultProfile) Render() image.Image {
  dc := gg.NewContext(600, 400)
  // BASE COLOR
  dc.SetHexColor("212121")
  dc.Clear()
  
  // BACKGROUND
  img := p.Generator.Toolbox.ReadImageFromURL(p.BackgroundURL, 600, 190)
  dc.DrawImage(img, 0, 0)
  rect := gg.NewContext(600, 190)
  rect.SetRGBA(0, 0, 0, 98)
  rect.Clear()
  dc.DrawImage(rect.Image(), 0, 0)
  
  // R RECT USERNAME
  dc.SetHexColor(p.FavColor)
  dc.DrawRoundedRectangle(180, 160, 170, 35, 5)
  dc.Fill()
  
  // WHITE CONTOURS
  dc.SetHexColor("ffffff")
  dc.SetLineWidth(6)
  dc.DrawLine(0, 193, 600, 193)
  dc.Stroke()
  
  /* FAV COLOR DEPENDENT OUTLINING */
  // AVATAR
  dc.SetHexColor(p.FavColor)
  dc.DrawCircle(90, 190, 74)
  dc.Fill()
  
  // LINE
  dc.SetHexColor(p.FavColor)
  dc.SetLineWidth(5)
  dc.DrawLine(0, 190, 600, 190)
  dc.Stroke()
  
  
  /* AVATAR DRAWING */
  avatarSize := 150
  x := 90
  y := 190
  
  // WHITE OUTLINE @ AVATAR
  dc.SetHexColor("ffffff")
  dc.DrawCircle(float64(x), float64(y), 70)
  dc.Fill()
  
  avatar := p.Generator.Toolbox.ReadImageFromURL(p.AvatarURL, avatarSize, avatarSize)
  dc.DrawCircle(float64(x), float64(y), 66)
  dc.Clip()
  dc.DrawImageAnchored(avatar, x, y, 0.5, 0.5)
  
  /* TEXT RENDERING SECTION */
  // USERNAME
  dc.SetHexColor(p.Generator.Toolbox.GetCompatibleFontColor(p.FavColor))
  
  dc.SavePNG("out.png")
  
  return dc.Image()
}