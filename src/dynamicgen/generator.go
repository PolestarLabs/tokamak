package dynamicgen

import (
	//"golang.org/x/image/font"
	//"image"
	"io/ioutil"
	//"os"
	"encoding/json"
	"strings"
	"tokamak/src/localizer"
)

type DynGenerator struct {
	Toolbox    Utils
	Templates  *localizer.Translator
	Blueprints map[string]Base
}

func New() DynGenerator {
	return DynGenerator{
		Toolbox:    NewUtil(),
		Templates:  localizer.New(),
		Blueprints: LoadBlueprints(),
	}
}

func LoadBlueprints() map[string]Base {
	files, err := ioutil.ReadDir("./blueprints")
	if err != nil {
		panic(err)
	}
	var ima []string
	for _, file := range files {
		ima = append(ima, strings.ReplaceAll(file.Name(), ".json", ""))
	}

	images := make(map[string]Base)
	for _, file := range ima {
		b, err := ioutil.ReadFile("./blueprints/" + file + ".json")
		if err != nil {
			panic(err)
		}

		var base Base

		err = json.Unmarshal(b, &base)
		if err != nil {
			panic(err)
		}

		images[file] = OptimizeBlueprint(base)
		b = []byte{}
	}

	return images
}
