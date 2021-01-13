package main
// utility for resizing images

import (
  "image"
  "image/png"
  _ "image/jpeg"
  "os"
)

import (
  "github.com/disintegration/imaging"
)

// X×Y
// Stickers: 100x100
// Emojis: 15x15
// Bgs: 600x190

func main () {
  path := "./assets/images/bgs/mcthah_red.png"
  x := 600
  y := 190
  save_to := "./assets/images/bgs/mctha_red.png"
  
  img_reader, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  defer img_reader.Close()
  
  img, _, err := image.Decode(img_reader)
  if err != nil {
    panic(err)
  }
  img = imaging.Fill(img, x, y, imaging.Center, imaging.Linear)
  
  outf, err := os.Create(save_to)
  if err != nil {
    panic(err)
  }
  defer outf.Close()
  
  err = png.Encode(outf, img)
  if err != nil {
    panic(err)
  } else {
    print("done")
  }
}