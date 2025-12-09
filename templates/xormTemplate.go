package _templates

const StructXORMTemplateContent = `{{range .Tables}}
// {{.Name | title}} represents the {{.Name}} table
type {{.Name | title}} struct {
{{range .Columns}}	{{.Name | title}} {{.Type}} ` + "`xorm:\"{{.Tag}}\"`" + ` // {{.Comment}}
{{end}}}
{{end}}
`

const XormTableNameTemplateContent = `{{range .Tables}}
// TableName specifies the table name for xorm
func (this *{{.Name | title}}) TableName() string {
	return "{{.Name}}"
}
{{end}}
`
