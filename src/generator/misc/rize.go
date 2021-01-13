package miscgenerator

import (
  "image"
  "tokamak/src/generator"
  "github.com/fogleman/gg"
)

type RizeData struct {
  Text string `json:"text" form:"text"`
}

func RenderRizeImage(g generator.Generator, p *RizeData) image.Image {
  dc := gg.NewContext(1192, 670)
  img := g.Toolbox.GetAsset("foundation/rize")
  dc.DrawImage(img, 0, 0)
  
  dc.LoadFontFace("../assets/fonts/Poppins/Poppins-Medium.ttf", 25)
  dc.SetHexColor("000000")
  
  dc.Rotate(0.11)
  g.Toolbox.DrawTextWrapped(dc, p.Text, 255, 93, 438, 840, 26)

  return dc.Image()
}