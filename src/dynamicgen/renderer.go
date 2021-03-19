package dynamicgen

import (
	"github.com/fogleman/gg"
	"image"
	"strings"
)

func (d *DynGenerator) Render(typ string, data map[string]interface{}) image.Image {
	if x, found := d.Blueprints[typ]; found {
		var ctx *gg.Context
		if x.DoNotGenerateContext {
			ctx = gg.NewContextForImage(d.Toolbox.GetAsset(strings.ReplaceAll(x.Entities[0].(map[string]interface{})["url"].(string), "assets://", "")))
		} else {
			ctx = gg.NewContext(x.Wh[0], x.Wh[1])
		}

		for i, entity := range x.Entities {
			if i == 0 && x.DoNotGenerateContext {
				continue
			}
			d.HandleEntity(ctx, entity.(map[string]interface{}), data)
		}

		return ctx.Image()
	} else {
		return nil
	}
}
