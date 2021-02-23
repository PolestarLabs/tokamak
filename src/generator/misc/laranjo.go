package miscgenerator

import (
	"github.com/fogleman/gg"
	"image"
	"tokamak/src/generator"
)

type LaranjoData struct {
	Text string `json:"text" form:"text"`
}

func RenderLaranjoImage(g generator.Generator, p *LaranjoData) image.Image {
	dc := gg.NewContext(685, 494)
	img := g.Toolbox.GetAsset("foundation/laranjo")
	dc.DrawImage(img, 0, 0)

	dc.LoadFontFace("../assets/fonts/Poppins/Poppins-Medium.ttf", 25)
	dc.SetHexColor("000000")

	g.Toolbox.DrawTextWrapped(dc, p.Text, 20, 30, 655, 800, 26)

	return dc.Image()
}
