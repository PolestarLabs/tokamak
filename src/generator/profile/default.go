package profilegenerator

import (
  "image"
  "tokamak/src/generator"
)

import "github.com/fogleman/gg"

type DefaultProfile struct {
  Generator generator.Generator
  AvatarURL, Background, Sticker string
  Married bool
  Name, AboutMe, FavColor, Money, PartnerName string
}

func (p DefaultProfile) Render() image.Image {
  dc := gg.NewContext(600, 400)
  // BASE COLOR
  dc.SetHexColor("212121")
  dc.Clear()
  
  dc.LoadFontFace("../assets/fonts/Poppins/Poppins-Medium.ttf", 12)
  
  // BACKGROUND
  img := p.Generator.Toolbox.GetAsset("bgs/"+ p.Background)
  dc.DrawImage(img, 0, 0)
  rect := gg.NewContext(600, 190)
  rect.SetRGBA(0, 0, 0, 98)
  rect.Clear()
  dc.DrawImage(rect.Image(), 0, 0)
  
  // R RECT USERNAME
  dc.SetHexColor(p.FavColor)
  dc.DrawRoundedRectangle(180, 160, 170, 35, 5)
  dc.Fill()
  
  /* STICKER DRAWING. */
  /* WE HAVE TO FIND OUT IF THE USER IS MARRIED. IF THEY ARE,
  WE NEED TO DRAW THE MONEY INDICATOR AND STICKER LOWER ON THE CARD (AFTER MARRIAGE STAT).*/
  
  // DEF POS (SINGLE)
  stickery := 40
  moneyy := 10
  
  if p.Married {
    moneyy += 30
    stickery += 30
    
    // R RECT MARRIAGE
    dc.SetRGBA(1, 1, 1, 170)
    dc.DrawRoundedRectangle(465, 10, 125, 25, 10)
    dc.Fill()
    
    // RING EMOJI / MARRIED
    img = p.Generator.Toolbox.GetAsset("emojis/ring")
    dc.DrawImage(img, 470, 15)
    
    dc.SetHexColor("000000")
    p.Generator.Toolbox.SafeDrawString(dc, p.PartnerName, 490, 26, 105)
  }
  
  // R RECT YEN
  dc.SetRGBA(1, 1, 1, 170)
  dc.DrawRoundedRectangle(490, float64(moneyy), 100, 25, 10)
  dc.Fill()
  
  // R RECT STICKER BLUR
  dc.SetRGBA(1, 1, 1, 170)
  dc.DrawRoundedRectangle(490, float64(stickery), 100, 100, 20)
  dc.Fill()
  // R RECT ABOUTME BLUR
  dc.SetRGBA(0, 0, 0, 180)
  dc.DrawRoundedRectangle(350, 225, 225, 150, 20)
  dc.Fill()
  // STICKER DRAWING
  img = p.Generator.Toolbox.GetAsset("stickers/"+ p.Sticker)
  dc.DrawImage(img, 490, stickery)
  
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
  
  // ABOUT ME
  dc.SetHexColor(p.FavColor)
  dc.DrawRoundedRectangle(365, 217, 110, 20, 10)
  dc.Fill()
  
  /* EMOJIS (15x15) */
  // YEN/MONEY
  img = p.Generator.Toolbox.GetAsset("emojis/money")
  dc.DrawImage(img, 497, moneyy + 3)
  
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
  dc.ResetClip()
  
  // YENS
  dc.SetHexColor("000000")
  dc.DrawString(p.Money, 519, float64(moneyy) + 16)
  
  // "ABOUT" ME" @ about ME
  img = p.Generator.Toolbox.GetAsset("emojis/woman_laptop")
  dc.DrawImage(img, 373, 219)
  
  dc.SetHexColor(p.Generator.Toolbox.GetCompatibleFontColor(p.FavColor))
  dc.DrawString("About me", 393, 231)
  
  dc.SetHexColor("ffffff")
  dc.DrawStringWrapped(p.AboutMe, 460, 275, 0.5, 0.5, 200, 1.5, gg.AlignLeft)
  
  /* TEXT RENDERING SECTION */
  dc.LoadFontFace("../assets/fonts/Poppins/Poppins-Bold.ttf", 23)
  
  // USERNAME
  // FIND OUT WHICH ONE IS MORE VISIBLE WITH THE CURRENT COLOR SCHEME: BLACK OR WHITE.
  dc.SetHexColor(p.Generator.Toolbox.GetCompatibleFontColor(p.FavColor))
  p.Generator.Toolbox.SafeDrawString(dc, p.Name, 185, 183, 160)
  
  dc.SavePNG("out.png")
  
  return dc.Image()
}