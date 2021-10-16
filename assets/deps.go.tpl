{{define "dependencies"}}
{{- /*gotype: github.com/raf924/bot-builder/internal/pkg/recipe.Recipe*/ -}}
package main
{{ range .Deps }}
//go:generate go mod download {{.Path}}@{{.Version}}
//go:generate go mod edit -require={{.Path}}@{{.Version}}
{{ end }}
//go:generate go mod tidy
import (
{{- /*gotype: github.com/raf924/bot-builder/internal/pkg/recipe.Recipe*/ -}}
{{ range .Deps }}
_ "{{.Path}}"
{{ end }}
)
{{end}}