package dynamicgen

import (
	"fmt"
	"github.com/fogleman/gg"
	ms "github.com/mitchellh/mapstructure"
	"strings"
)

func (gen *DynGenerator) HandleEntity(ctx *gg.Context, d map[string]interface{}, data map[string]interface{}) {

	switch d["type"].(string) {
	case "image":
		var im ImageEntity
		err := ms.Decode(d, &im)
		if err != nil {
			fmt.Println("failed to cast image entity from map[string]interface{} to ImageEntity")
			break
		}

		if strings.HasPrefix(im.Url, "assets://") {
			img := gen.Toolbox.GetAsset(strings.ReplaceAll(im.Url, "assets://", ""))
			ctx.DrawImage(img, im.Xy[0], im.Xy[1])
		} else {
			fmt.Println("xereca")
		}
	case "rotate":
		ctx.Rotate(d["value"].(float64))
	case "text":
		var tx TextEntity
		err := ms.Decode(d, &tx)
		if err != nil {
			fmt.Println(err)
			break
		}
		properc := gen.Templates.Execute(tx.Content, data)
		if tx.Font.Face != "" {
			ctx.LoadFontFace("./assets/fonts/"+tx.Font.Face+".ttf", tx.Font.Size)
		}

		if tx.Color != "" {
			ctx.SetHexColor(tx.Color)
		}

		if tx.MaxWidth == 0.0 && tx.MaxHeight == 0.0 { // if there are no limits set for this text
			ctx.DrawString(properc, tx.Xy[0], tx.Xy[1])
		} else if tx.MaxWidth != 0.0 && tx.MaxHeight == 0.0 { // if there's a horizontal limit but no vertical
			gen.Toolbox.SafeDrawString(ctx, properc, tx.Xy[0], tx.Xy[1], tx.MaxWidth)
		} else {
			gen.Toolbox.DrawTextWrapped(ctx, properc, tx.Xy[0], tx.Xy[1], tx.MaxWidth, tx.MaxHeight, tx.Spacing)
		}
	default:
		fmt.Printf("%s entity not implemented\n", d["type"].(string))
	}
}
