package profilegenerator

import (
	"image"
	"strings"
	_ "strings"
	"tokamak/src/generator"
)

import "github.com/fogleman/gg"

type ProfileData struct {
	AvatarURL   string `json:"avatarUrl" form:"avatarUrl"`
	Background  string `json:"bgId" form:"bgId"`
	Sticker     string `json:"stickerId" form:"stickerId"`
	Married     bool   `json:"married" form:"married"`
	Name        string `json:"name" form:"name"`
	AboutMe     string `json:"aboutMe" form:"aboutMe"`
	FavColor    string `json:"favColor" form:"favColor"`
	Money       string `json:"money" form:"money"`
	Type        string `json:"type" form:"type"`
	PartnerName string `json:"partnerName" form:"partnerName"`
}

func RenderDefaultProfile(g generator.Generator, p *ProfileData) image.Image {
	dc := gg.NewContext(600, 400)
	// BASE COLOR
	dc.SetHexColor("#212121")
	dc.Clear()

	_ = dc.LoadFontFace("./assets/fonts/Poppins/Poppins-Medium.ttf", 15)

	// BACKGROUND
	img := g.Toolbox.GetAsset("bgs/" + p.Background)
	dc.DrawImage(img, 0, 0)
	rect := gg.NewContext(600, 190)
	rect.SetRGBA(0, 0, 0, 98)
	rect.Clear()
	dc.DrawImage(rect.Image(), 0, 0)

	// FAV COLOR STROKE
	space := float64(len(strings.TrimSpace(p.Name))) * 16 //FIXME PORRA DE INFERO NÃO FICA DO JEITO CERTO
	if space >= 290 {
		space = 280
	}
	dc.SetHexColor(p.FavColor)
	dc.DrawRoundedRectangle(182, 147, space+21, 48, 14) // <-
	dc.Fill()

	// LINE
	dc.SetHexColor(p.FavColor)
	dc.SetLineWidth(5)
	dc.DrawLine(0, 190, 600, 190)
	dc.Stroke()

	// R RECT USERNAME
	dc.SetHexColor("#FFFFFF")
	dc.DrawRoundedRectangle(185, 150, space, 45, 14) // <-
	dc.Fill()

	// WHITE LINE SO THE ROUNDED RECTANGLE CAN LOOK STRAIGHT
	dc.SetHexColor("#FFFFFF")
	dc.DrawLine(187, 172, 187, 191)
	dc.SetLineWidth(8)
	dc.DrawLine(335, 172, 335, 191)
	dc.Stroke()

	// FAV COLOR LINE SO THE ROUNDED RECTANGLE CAN LOOK STRAIGHT
	dc.SetHexColor(p.FavColor)
	dc.SetLineWidth(3)
	dc.DrawLine(183.2, 172, 183.2, 188)
	dc.DrawLine(341.9, 172, 341.9, 188)
	dc.Stroke()

	/* STICKER DRAWING */

	/* WE HAVE TO FIND OUT IF THE USER IS MARRIED. IF THEY ARE,
	WE NEED TO DRAW THE MONEY INDICATOR AND STICKER LOWER ON THE CARD (AFTER MARRIAGE STAT).*/

	// DEF POS (SINGLE)
	stickerY := 40
	moneyY := 10

	if p.Married {
		moneyY += 30
		stickerY += 30

		// R RECT MARRIAGE
		dc.SetHexColor("#878787")
		dc.DrawRoundedRectangle(465, 10, 124, 25, 10)
		dc.Fill()

		// RING EMOJI / MARRIED
		img = g.Toolbox.GetAsset("emojis/ring")
		dc.DrawImage(img, 470, 15)

		dc.SetHexColor("#FFFFFF")
		g.Toolbox.SafeDrawString(dc, p.PartnerName, 490, 27, 105)
	}

	// R RECT YEN
	dc.SetHexColor("#878787")
	dc.DrawRoundedRectangle(465, float64(moneyY), 124, 25, 8)
	dc.Fill()

	// R RECT STICKER BLUR
	dc.SetHexColor("#878787")
	dc.DrawRoundedRectangle(490, float64(stickerY), 100, 100, 15)
	dc.Fill()

	// R RECT ABOUT ME BLUR
	dc.SetRGBA(0, 0, 0, 180)
	dc.DrawRoundedRectangle(350, 225, 225, 150, 20)
	dc.Fill()

	// STICKER DRAWING
	img = g.Toolbox.GetAsset("stickers/" + p.Sticker)
	dc.DrawImage(img, 490, stickerY)

	// WHITE CONTOURS
	dc.SetHexColor("#ffffff")
	dc.SetLineWidth(4)
	dc.DrawLine(0, 193, 600, 193)
	dc.Stroke()

	/* FAV COLOR DEPENDENT OUTLINING */
	// AVATAR
	dc.SetHexColor(p.FavColor)
	dc.DrawCircle(90, 190, 74)
	dc.Fill()

	// REPS

	// ABOUT ME ROUNDED RECTANGLE
	dc.SetHexColor(p.FavColor)
	dc.DrawRoundedRectangle(365, 217, 110, 20, 10)
	dc.Fill()

	// WHITE CIRCLE @ ABOUT ME
	dc.SetHexColor("#FFFFFF")
	dc.DrawCircle(375, 227, 14)
	dc.Fill()

	/* EMOJIS (15x15) */
	// YEN/MONEY
	img = g.Toolbox.GetAsset("emojis/money")
	dc.DrawImage(img, 470, moneyY+3)

	/* AVATAR DRAWING */
	avatarSize := 150
	x := 90
	y := 190

	// WHITE OUTLINE @ AVATAR
	dc.SetHexColor("#ffffff")
	dc.DrawCircle(float64(x), float64(y), 70)
	dc.Fill()

	avatar := g.Toolbox.ReadImageFromURL(p.AvatarURL, avatarSize, avatarSize)
	dc.DrawCircle(float64(x), float64(y), 66)
	dc.Clip()
	dc.DrawImageAnchored(avatar, x, y, 0.5, 0.5)
	dc.ResetClip()

	// YENS
	dc.SetHexColor("#FFFFFF")
	dc.DrawString(p.Money, 519, float64(moneyY)+18)

	// "ABOUT" ME" @ about ME
	img = g.Toolbox.GetAsset("emojis/woman_laptop")
	dc.DrawImage(img, 365, 218)

	dc.SetHexColor(g.Toolbox.GetCompatibleFontColor(p.FavColor))
	dc.DrawString("About me", 393, 232)

	dc.SetHexColor("#ffffff")
	dc.DrawStringWrapped(p.AboutMe, 460, 265, 0.5, 0.5, 200, 1.5, gg.AlignLeft)

	/* TEXT RENDERING SECTION */
	_ = dc.LoadFontFace("./assets/fonts/Poppins/Poppins-Bold.ttf", 23)

	// USERNAME
	// FIND OUT WHICH ONE IS MORE VISIBLE WITH THE CURRENT COLOR SCHEME: BLACK OR WHITE.
	dc.SetHexColor("#000000")
	_ = dc.LoadFontFace("./assets/fonts/Poppins/Poppins-Light.ttf", 27)
	g.Toolbox.SafeDrawString(dc, p.Name, 192, 180, 290)

	return dc.Image()
}
