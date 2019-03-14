package templates

import "text/template"

var (
	Generator = template.Must(template.New("generator").Parse(`{{- /*gotype: github.com/Adphi/go-mobile-collection.GenerateTemplateData*/ -}}
// code generated by go-mobile-collection. DO NOT EDIT.
// WARNING - These collections are not thread-safe

package {{.Package}}

import (
"encoding/json"
"github.com/pkg/errors"
)

{{ range .Types }}
{{ if .TypeNamed }}{{ template "typeNamed" . }}{{ else }}{{ template "embed" . }}{{ end }}

{{ end }}`))
)
