package utils

import (
  "net/http"
  "math"
  "image"
  "errors"
  "image/color"
  _ "image/jpeg"
  _ "image/png"
  "golang.org/x/image/font"
)

import (
  "github.com/fogleman/gg"
  "github.com/disintegration/imaging"
  //"github.com/golang/freetype/truetype"
)

type Utils struct {
  default_image image.Image
  fontCache []Font
}

type Font struct {
  font font.Face
  path string
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
  imagem = imaging.Fill(imagem, x, y, imaging.Center, imaging.Linear)
  return imagem
}

func (util *Utils) GetColorLuminance(color color.RGBA) float64 {
  return float64(float64(0.299) * float64(color.R) + float64(0.587) * float64(color.G) + float64(0.114) * float64(color.B))
}

/*func (util Utils) GetFontFaceOrCreate(path string) font.Face {
  append(util.fontCache, face)
}*/

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

func NewUtil () Utils {
  def := gg.NewContext(1000, 1000)
  def.SetRGB(0.2, 0.2, 0.2)
  def.Clear()
  
  return Utils {
    default_image: def.Image(),
  }
}