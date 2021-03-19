package localizer

import (
	"bytes"
	"fmt"
	"github.com/leekchan/gtf"
	"strings"
	"text/template"
)

type Translator struct {
	Temp *template.Template
}

func (t *Translator) T(lang, key string, data interface{}) string {
	if lang == "" {
		lang = "pt-br"
	}
	final := new(bytes.Buffer)
	err := t.Temp.ExecuteTemplate(final, lang+"_"+key, data)

	if err != nil {
		return key
	}
	return strings.ReplaceAll(final.String(), "\n", "")
}

func (t *Translator) Execute(tem string, data interface{}) string {
	temp, err := t.Temp.Parse(tem)
	if err != nil {
		fmt.Println(err)
		return "?"
	}

	rst := new(bytes.Buffer)
	err = temp.Execute(rst, data)
	if err != nil {
		fmt.Println(err)
		return "?"
	}

	return rst.String()
}

func New() *Translator {
	t := Translator{}
	a := template.Must(
		template.New("hm").
			Funcs(gtf.GtfTextFuncMap).
			Funcs(template.FuncMap{
				"t": t.T,
			}).
			ParseGlob("./locales/*"))
	t.Temp = a

	return &t
}
