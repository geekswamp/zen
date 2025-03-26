package template

import (
	"fmt"
	"os"
	"reflect"
	"text/template"

	"github.com/geekswamp/zen/cmd/genz/internal/format"
	"github.com/geekswamp/zen/cmd/genz/internal/mod"
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

func ParseModel(model any) (names []string, types []reflect.Type) {
	t := reflect.TypeOf(model)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, nil
	}

	var parseFields func(t reflect.Type)

	parseFields = func(t reflect.Type) {
		for i := range t.NumField() {
			field := t.Field(i)

			if field.Anonymous {
				parseFields(field.Type)
				continue
			}

			names = append(names, field.Name)
			types = append(types, field.Type)
		}
	}

	parseFields(t)

	return names, types
}
