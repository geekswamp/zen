package template

import (
	"fmt"
	"os"
	"text/template"

	"github.com/geekswamp/genz/internal/format"
	"github.com/geekswamp/genz/internal/mod"
)

func (m Make) Parse(file *os.File) error {
	tmplFile := fmt.Sprintf("%s.tmpl", m.FileType)

	funcMap := template.FuncMap{
		"ToCamelCase": func(text string) string {
			return format.ToCamelCase(text)
		},
		"ToPascalCase": func(text string) string {
			return format.ToPascalCase(text)
		},
	}

	t, err := template.New(tmplFile).Funcs(funcMap).ParseFS(TemplateFile, tmplFile)
	if err != nil {
		return err
	}

	modName, err := mod.GetModuleName()
	if err != nil {
		return err
	}

	data := map[string]any{
		"Module":     modName,
		"StructName": m.FeatureName,
	}

	if err := t.Execute(file, data); err != nil {
		return err
	}

	return nil
}
