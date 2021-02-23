package generator

import (
	"golang.org/x/image/font"
	"image"
	"tokamak/src/utils"
)

type Generator struct {
	Toolbox    utils.Utils
	FontCache  []Font
	ImageCache []ImageAsset
}

type Font struct {
	path string
	face font.Face
}

type ImageAsset struct {
	path string
	img  image.Image
}

func NewGenerator() Generator {
	return Generator{
		FontCache:  []Font{},
		ImageCache: []ImageAsset{},
		Toolbox:    utils.NewUtil(),
	}
}
