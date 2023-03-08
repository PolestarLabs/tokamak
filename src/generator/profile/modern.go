package profilegenerator

import (
	"image"
	"strings"
	"time"
	"tokamak/src/generator"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)


func RenderModernProfile(g generator.Generator, p *ProfileData) image.Image {
	dc := gg.NewContext(1100, 720)
	
	/* BACKGROUND */
	img := g.Toolbox.GetAsset("bgs/" + p.Background)


	g.AddFontInCtx(dc, "../assets/fonts/Ghost/iknowaghost.ttf", 36) // Load First Font

	/* Married	*/
	if p.Married {

		/* Load Bar Float */
		img := g.Toolbox.GetAsset("/template/married")
		dc.DrawImage(img, 0, 0)

		/* Partner */
		dc.SetHexColor(g.Toolbox.GetCompatibleFontColor("#ffff"))
		dc.DrawString(p.PartnerName, float64(340), float64(108))
		
	}

	dc.LoadFontFace("../assets/fonts/Ghost/iknowaghost.ttf", 50)




	// Circle Profile
	dc.SetHexColor(p.FavColor)
	dc.DrawRoundedRectangle(180, 160, 170, 35, 5)
	dc.Fill()

		
	img = g.Toolbox.GetAsset("template/profile")
	dc.DrawImage(img, 0, 0)



	/* Sticker */

	img = g.Toolbox.GetAsset("stickers/modern/" + p.Sticker)
	dc.DrawImage(img, 260, 212)


	/* Tippy icon	*/
	img = g.Toolbox.GetAsset("template/hey-tippy")
	dc.DrawImage(img, 0, 0)
	imaging.AdjustContrast(img, 0)

	/* 	 Chino Kafuu   */
	img = g.Toolbox.GetAsset("template/part-1")
	dc.DrawImage(img, 0, 0)
	
	

	/* AVATAR DRAWING */
	avatarSize := 290
	circleSize := float64(129)
	x := 168
	y := 181

	/* Outline of avatar	*/
	dc.SetHexColor("ffffff")
	dc.DrawCircle(float64(x), float64(y), 70)
	dc.Fill()

	dc.SetHexColor(p.FavColor)
	dc.DrawCircle(float64(x), float64(y), 140)
	dc.Fill()

	avatar := g.Toolbox.ReadImageFromURL(p.AvatarURL, avatarSize, avatarSize)
	dc.DrawCircle(float64(x), float64(y), circleSize)
	dc.Clip()
	dc.DrawImageAnchored(avatar, x, y, 0.5, 0.5)
	dc.ResetClip()


	/* NickName	*/
	dc.SetHexColor(g.Toolbox.GetCompatibleFontColor("#ffff"))
	dc.DrawString(p.Name, float64(314), float64(178))



	/* Yens	*/ 
	g.AddFontInCtx(dc, "../assets/fonts/Ghost/iknowaghost.ttf", 30) // Default is 40px
	g.Toolbox.DrawTextWrapped(dc, p.Money, 140, 400, 208, 408, 13)

	/* About Me	*/
	dc.SetHexColor("#1d1c1d")
	g.Toolbox.DrawTextWrapped(dc, p.AboutMe, 590, 256, 360, 250, 50)



	/* Date	*/
	dc.SetHexColor("#1d1c1d")
	data := time.Now()
	dc.DrawString(strings.Replace(data.Format("02/Jan/2006"), "\\/{d:a}", " ", 0), float64(899), float64(178)) // Please don't ask me why I did this.

	
	

	/* BADGES */
	bx := 115.0
	by := 515.0
	badgesizex := 35.0
	badgesizey := 30.0
	badgespacing := 38.0
	spacebtwedge := 10.0
	recsizex := 850.0

	recsizey := badgesizey*2 + badgespacing


	cxpos := bx + spacebtwedge
	cypos := by + spacebtwedge

	nb := 0
	for _, b := range p.Badges {
		nb++
		if nb == 100 {
		} else {
			if b != "" {
				dc.DrawImage(g.Toolbox.GetAsset("badges/profile_2/"+b), int(cxpos), int(cypos))
			}
	
			cxpos = cxpos + badgesizex + badgespacing
			if cxpos > bx+spacebtwedge+recsizex {
				cypos = cypos + badgesizey + badgespacing
				cxpos = bx + spacebtwedge
				if cypos > by+spacebtwedge+recsizey {
					break // if we run out of space, break the loop
				}
			}
		}
	
	}




	return dc.Image()
}
