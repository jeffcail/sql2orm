package _templates

// 纯内容模板：只生成 struct 主体
const StructGORMTemplateContent = `{{range .Tables}}
// {{.Name | title}} represents the {{.Name}} table
type {{.Name | title}} struct {
{{range .Columns}}	{{.Name | title}} {{.Type}} ` + "`gorm:\"{{.Tag}}\"`" + ` // {{.Comment}}
{{end}}}
{{end}}
`

// 纯内容模板：只生成 TableName 方法
const GormTableNameTemplateContent = `{{range .Tables}}
// TableName overrides the table name used by GORM
func ({{.Name | title}}) TableName() string {
	return "{{.Name}}"
}
{{end}}
`
