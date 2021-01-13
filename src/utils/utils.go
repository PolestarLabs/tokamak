package utils

import (
  "net/http"
  "math"
  "image"
  "errors"
  "image/color"
  _ "image/jpeg"
  _ "image/png"
  "os"
  "unicode/utf8"
)

import (
  "github.com/fogleman/gg"
  "github.com/disintegration/imaging"
)

type Utils struct {
  default_image image.Image
  asset_cache map[string]image.Image
}

func (util *Utils) SafeDrawStringAnchored(ctx *gg.Context, text string, x, y, w, ax, ay float64) {
  ctext := text
  
  for ok := getWidth(ctx.MeasureString(ctext)) > w; ok; ok = getWidth(ctx.MeasureString(ctext)) > w {
    ctext = util.TrimLastChar(ctext)
  }
  
  if getWidth(ctx.MeasureString(text)) > w {
    for z := 0; z < 3; z++ {
      ctext = util.TrimLastChar(ctext)
    }
    ctext = ctext + "..."
  }
  ctx.DrawStringAnchored(ctext, x, y, ax, ay)
}

func (util *Utils) SafeDrawString(ctx *gg.Context, text string, x, y, w float64) {
  ctext := text
  
  for ok := getWidth(ctx.MeasureString(ctext)) > w; ok; ok = getWidth(ctx.MeasureString(ctext)) > w {
    ctext = util.TrimLastChar(ctext)
  }
  
  if getWidth(ctx.MeasureString(text)) > w {
    for z := 0; z < 3; z++ {
      ctext = util.TrimLastChar(ctext)
    }
    ctext = ctext + "..."
  }
  ctx.DrawString(ctext, x, y)
}

func (util Utils) TrimLastChar(s string) string {
  r, size := utf8.DecodeLastRuneInString(s)
  if r == utf8.RuneError && (size == 0 || size == 1) {
    size = 0
  }
  return s[:len(s)-size]
}

func getWidth(w, h float64) float64 {
  return w
}

func (util *Utils) GetAsset(path string) image.Image {
  v, ok := util.asset_cache[path]
  if ok == true {
    return v
  }
  img_reader, err := os.Open("../assets/images/"+ path +".png")
  if err != nil {
    panic(err)
  }
  defer img_reader.Close()
  
  img, _, err := image.Decode(img_reader)
  if err != nil {
    panic(err)
  }
  
  util.asset_cache[path] = img
  return img
}

func (util *Utils) ReadImageFromURL(url string, x, y int) image.Image {
  var imagem image.Image = nil
  
  res, err := http.Get(url)
  if err != nil {
    panic(err)
    imagem = util.default_image
  }
  defer res.Body.Close()
  
  if imagem == nil {
    img, _, err := image.Decode(res.Body)
    if err != nil {
      panic(err)
      imagem = util.default_image
    } else {
      imagem = img
    }
  }
  imagem = imaging.Fill(imagem, x, y, imaging.Center, imaging.NearestNeighbor)
  return imagem
}

func (util *Utils) GetColorLuminance(color color.RGBA) float64 {
  return float64(float64(0.299) * float64(color.R) + float64(0.587) * float64(color.G) + float64(0.114) * float64(color.B))
}

func (util *Utils) GetCompatibleFontColor(hex_color string) string {
  c, err := util.ParseHexColor(hex_color)
  if err != nil {
    c = color.RGBA { R: 0, G: 0, B: 0, A: 0xff }
  }
  
  if math.Abs(util.GetColorLuminance(c) - util.GetColorLuminance(color.RGBA { R: 0, G: 0, B: 0, A: 0xff })) >= 128.0 {
    return "000000"
  } else {
    return "ffffff"
  }
}

var errInvalidFormat = errors.New("Invalid HEX color format!")

func (util *Utils) ParseHexColor(s string) (c color.RGBA, err error) {
  c.A = 0xff
  
  hexToByte := func(b byte) byte {
    switch {
    case b >= '0' && b <= '9':
      return b - '0'
    case b >= 'a' && b <= 'f':
      return b - 'a' + 10
    case b >= 'A' && b <= 'F':
      return b - 'A' + 10
    }
    
    err = errInvalidFormat
    return 0
  }

  switch len(s) {
  case 6:
    c.R = hexToByte(s[0])<<4 + hexToByte(s[1])
    c.G = hexToByte(s[2])<<4 + hexToByte(s[3])
    c.B = hexToByte(s[4])<<4 + hexToByte(s[5])
  case 3:
    c.R = hexToByte(s[0]) * 17
    c.G = hexToByte(s[1]) * 17
    c.B = hexToByte(s[2]) * 17
  default:
    err = errInvalidFormat
  }
  
  return
}

func canFitHeightWise(ctx *gg.Context, lines []string, maxHeight, spacing int) bool {
  sum := 0
  for _, text := range lines {
    _, h := ctx.MeasureString(text)
    sum += int(h) + spacing
  }
  return sum < maxHeight
}

func (util *Utils) DrawTextWrapped(ctx *gg.Context, s string, x, y, width, height, spacing int) {
  lines := ctx.WordWrap(s, float64(width))
  var tbd []string

  for len(lines) > 0 && canFitHeightWise(ctx, append(tbd, lines[0]), height, spacing) {
    tbd = append(tbd, lines[0])
    lines = lines[1:]
  }
  
  currentY := y
  for _, text := range tbd {
    ctx.DrawString(text, float64(x), float64(currentY))
    currentY += spacing
  }
}

func NewUtil () Utils {
  def := gg.NewContext(1000, 1000)
  def.SetRGB(0.2, 0.2, 0.2)
  def.Clear()
  
  return Utils {
    default_image: def.Image(),
    asset_cache: make(map[string]image.Image),
  }
}