package _templates

import (
	"strings"
	"text/template"

	"github.com/jeffcail/sql2orm/utils"
)

type TemplateSet struct {
	Struct     *template.Template
	TableName  *template.Template
	ImportTime bool
}

func title(s string) string {
	if s == "" {
		return ""
	}
	// 下划线转大驼峰
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
		}
	}
	return strings.Join(parts, "")
}

func getTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"Mapper": func(s string) string { return strings.Title(s) },
		"Type":   utils.SqlTypeToGoType,
		"Tag":    func(table utils.Table, col utils.Column) string { return col.Tag },
		"lower":  utils.ToUpperCamelCase,
		"title":  title,
	}
}

func getTemplateFuncstemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"title": title,
	}
}

func ParseGORM() (*TemplateSet, error) {
	structTmpl, _ := template.New("struct").Funcs(getTemplateFuncs()).Parse(StructGORMTemplateContent)
	tableNameTmpl, _ := template.New("tableName").Funcs(getTemplateFuncs()).Parse(GormTableNameTemplateContent)
	return &TemplateSet{structTmpl, tableNameTmpl, true}, nil
}

func ParseXORM() (*TemplateSet, error) {
	structTmpl, _ := template.New("struct").Funcs(getTemplateFuncs()).Parse(StructXORMTemplateContent)
	tableNameTmpl, _ := template.New("tableName").Funcs(getTemplateFuncs()).Parse(XormTableNameTemplateContent)
	return &TemplateSet{structTmpl, tableNameTmpl, true}, nil
}
