package main

// utility for resizing images

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

import (
	"github.com/disintegration/imaging"
)

// X×Y
// Stickers: 100x100
// Emojis: 15x15
// Bgs: 600x190

func main() {
	path := "./menina.png"
	x := 18
	y := 18
	save_to := "./assets/images/emojis/woman_laptop.png"

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
