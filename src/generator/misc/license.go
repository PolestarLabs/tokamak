package miscgenerator

import (
	"github.com/fogleman/gg"
	"image"
	"tokamak/src/generator"
)

type LicenseData struct {
	Text      string `json:"text" form:"text"`
	Name      string `json:"name" form:"name"`
	AvatarURL string `json:"avatarUrl" form:"avatarUrl"`
	HexColor  string `json:"hexColor" form:"hexColor"`
}

func RenderLicenseImage(g generator.Generator, p *LicenseData) image.Image {
	dc := gg.NewContext(1150, 893)
	img := g.Toolbox.GetAsset("foundation/license")
	avatar := g.Toolbox.ReadImageFromURL(p.AvatarURL, 393, 393)

	dc.Rotate(-0.2) // Math.Pi * 2 / -30

	// Rectangle //
	dc.SetHexColor(p.HexColor)
	dc.DrawRoundedRectangle(-10, 250, 860, 550, 50)
	dc.Fill()
	// --------- //

	dc.DrawImage(img, -10, 300)
	dc.DrawImage(avatar, 428, 300)

	dc.LoadFontFace("../assets/fonts/Poppins/Poppins-Bold.ttf", 45)
	dc.SetHexColor("000000")
	g.Toolbox.SafeDrawString(dc, p.Name, 5, 380, 445)

	dc.SetHexColor(g.Toolbox.GetCompatibleFontColor(p.HexColor))
	g.Toolbox.SafeDrawStringAnchored(dc, p.Text, 400, 740, 830, 0.5, 0.5)

	return dc.Image()
}
