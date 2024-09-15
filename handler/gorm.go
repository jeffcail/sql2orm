package handler

import (
	_templates "github.com/jeffcail/sql2orm/templates"
	"github.com/jeffcail/sql2orm/utils"
	"strings"
	"text/template"
)

func parseGorm() (*template.Template, *template.Template, error) {
	structTmpl, err := template.New("struct").Funcs(template.FuncMap{
		"Mapper": func(s string) string { return strings.Title(s) },
		"Type":   utils.SqlTypeToGoType,
		"Tag":    func(table utils.Table, col utils.Column) string { return col.Tag },
		"lower":  utils.ToUpperCamelCase,
	}).Parse(_templates.StructGORMTemplateContent)
	if err != nil {
		return nil, nil, err
	}

	tableNameTmpl, err := template.New("tableName").Funcs(template.FuncMap{
		"Mapper": func(s string) string { return strings.Title(s) },
	}).Parse(_templates.GormTableNameTemplateContent)
	if err != nil {
		return nil, nil, err
	}
	return structTmpl, tableNameTmpl, nil

}
