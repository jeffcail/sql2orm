package _templates

const StructTemplateContent = `
package models

import (
	"time"
)

{{- range .Tables }}
// {{ .Name }} represents a row in the '{{ .TableName }}' table.
type {{ .Name }} struct {
{{- $table := . }}
{{- range .Columns }}
	{{ .Name | lower }}	{{ .Type }} ` + "`xorm:\"{{ .Tag }}\"`" + ` // {{ .Comment }}
{{- end }}
}
{{- end }}
`

const TableNameTemplateContent = `
{{- range .Tables }}
// TableName sets the insert table name for this struct type
func (this *{{ .Name }}) TableName() string {
	return "{{ .TableName }}"
}
{{- end }}
`
