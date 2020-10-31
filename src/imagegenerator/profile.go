package imagegenerator

import (
  "image"
  "tokamak/src/utils"
)

import "github.com/fogleman/gg"

type ProfileGenerator struct {
  Toolbox utils.Utils
}

type ProfileData struct {
  AvatarURL, BackgroundURL string
  Name, AboutMe, FavColor, HighestRole string
}

func (p ProfileGenerator) Render(data ProfileData) image.Image {
  dc := gg.NewContext(600, 400)
  // BASE COLOR
  dc.SetHexColor("212121")
  dc.Clear()
  
  // BACKGROUND
  img := p.Toolbox.ReadImageFromURL(data.BackgroundURL, 600, 190)
  dc.DrawImage(img, 0, 0)
  rect := gg.NewContext(600, 190)
  rect.SetRGBA(0, 0, 0, 98)
  rect.Clear()
  dc.DrawImage(rect.Image(), 0, 0)
  
  // R RECT USERNAME
  dc.SetHexColor(data.FavColor)
  dc.DrawRoundedRectangle(180, 160, 170, 35, 5)
  dc.Fill()
  
  // WHITE CONTOURS
  dc.SetHexColor("ffffff")
  dc.SetLineWidth(6)
  dc.DrawLine(0, 193, 600, 193)
  dc.Stroke()
  
  /* FAV COLOR DEPENDENT OUTLINING */
  // AVATAR
  dc.SetHexColor(data.FavColor)
  dc.DrawCircle(90, 190, 74)
  dc.Fill()
  
  // LINE
  dc.SetHexColor(data.FavColor)
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
  
  avatar := p.Toolbox.ReadImageFromURL(data.AvatarURL, avatarSize, avatarSize)
  dc.DrawCircle(float64(x), float64(y), 66)
  dc.Clip()
  dc.DrawImageAnchored(avatar, x, y, 0.5, 0.5)
  
  /* TEXT RENDERING SECTION */
  // Username
  dc.SetHexColor(p.Toolbox.GetCompatibleFontColor(data.FavColor))
  
  dc.SavePNG("out.png")
  
  return dc.Image()
}