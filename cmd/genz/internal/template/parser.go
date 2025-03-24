package template

import (
	"fmt"
	"os"
	"text/template"

	"github.com/geekswamp/genz/internal/format"
)

func (m Make) Parse(file *os.File) error {
	tmplFile := fmt.Sprintf("%s.tmpl", m.fileType)

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

	data := map[string]any{
		"StructName": m.featureName,
	}

	if err := t.Execute(file, data); err != nil {
		return err
	}

	return nil
}
