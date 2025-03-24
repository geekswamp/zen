package format

import "github.com/iancoleman/strcase"

func ToCamelCase(text string) string {
	return strcase.ToLowerCamel(text)
}

func ToPascalCase(text string) string {
	return strcase.ToCamel(text)
}
