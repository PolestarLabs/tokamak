package dynamicgen

import (
	ms "github.com/mitchellh/mapstructure"
	"strings"
)

// optimizes blueprints
// TODO: check if a text entity requires an already loaded font face. if it is, drop the redundant font object
func OptimizeBlueprint(b Base) Base {
	if b.Entities[0].(map[string]interface{})["type"] == "image" {
		var img ImageEntity
		_ = ms.Decode(b.Entities[0].(map[string]interface{}), &img)

		if strings.HasPrefix(img.Url, "assets://") && img.Xy == [2]int{0, 0} {
			b.DoNotGenerateContext = true
		}
	}

	return b
}
