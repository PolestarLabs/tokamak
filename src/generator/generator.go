package generator

import (
	"fmt"
	"image"
	"io/ioutil"
	"regexp"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"tokamak/src/utils"
)

var GeneratorVar = Generator{
	FontCache:  []*Font{},
	ImageCache: []*ImageAsset{},
	Toolbox:    utils.NewUtil(),
}

type Generator struct {
	Toolbox    utils.Utils
	FontCache  []*Font
	ImageCache []*ImageAsset
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
	return GeneratorVar
}

func (generator *Generator) AddFontInCtx(ctx *gg.Context, fontPath string, size float64) (*gg.Context, error) {
	regex := regexp.MustCompile("\\.\\.")
	path := regex.ReplaceAllString(fontPath, "")
	fontIsCached := false

	for _, value := range generator.FontCache {
		if value.path == fmt.Sprintf(path, size) {
			ctx.SetFontFace(value.face)
			fontIsCached = true
			break
		}
	}
	if fontIsCached {
		return ctx, nil
	}
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return ctx, err
	}

	fontObj, err := truetype.Parse(fontBytes)
	if err != nil {
		return ctx, err
	}

	face := truetype.NewFace(fontObj, &truetype.Options{
		Size:              size,
		Hinting:           font.HintingFull,
		GlyphCacheEntries: 512,
	})

	generator.FontCache = append(generator.FontCache, &Font{
		path: fmt.Sprintf(path, size),
		face: face,
	})

	ctx.SetFontFace(face)

	return ctx, nil
}
