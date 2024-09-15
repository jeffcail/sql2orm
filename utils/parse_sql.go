package utils

import (
	"fmt"
	"strings"
)

type Column struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type Table struct {
	Name      string
	TableName string
	Columns   []Column
}

func ParseSQLTable(sql string) (Table, error) {
	var table Table

	// Remove surrounding whitespaces
	sql = strings.TrimSpace(sql)

	// Check if the SQL statement starts with "CREATE TABLE"
	if !strings.HasPrefix(strings.ToUpper(sql), "CREATE TABLE") {
		return table, fmt.Errorf("invalid SQL create table statement")
	}

	// Extract table name
	tableStart := strings.Index(sql, "(")
	if tableStart == -1 {
		return table, fmt.Errorf("invalid SQL create table statement")
	}
	tableDef := sql[:tableStart]
	tableDef = strings.TrimSpace(tableDef)
	tableDefParts := strings.Fields(tableDef)
	if len(tableDefParts) < 3 {
		return table, fmt.Errorf("invalid SQL create table statement")
	}
	table.Name = ToUpperCamelCase(strings.Title(strings.Trim(tableDefParts[2], "`")))
	table.TableName = strings.Trim(tableDefParts[2], "`")

	// Extract columns definitions
	columnsDef := sql[tableStart+1 : len(sql)-1]

	columns := strings.Split(columnsDef, ",")
	for _, column := range columns {
		column = strings.TrimSpace(column)
		columnParts := strings.Fields(column)
		if len(columnParts) < 2 {
			continue
		}
		columnName := strings.Title(strings.Trim(columnParts[0], "`"))
		columnType := columnParts[1]

		comment := ""
		commentStart := strings.Index(column, "COMMENT")
		if commentStart != -1 {
			commentPart := column[commentStart:]
			commentSplit := strings.SplitN(commentPart, "'", 2)
			if len(commentSplit) > 1 {
				comment = strings.Trim(commentSplit[1], "'")
			}
		}

		var tagStr string
		if columnName == "Id" {
			tagStr = CompactStr("'", strings.ToLower(columnName), "'", " pk autoincr")
		} else {
			tagStr = CompactStr("'", strings.ToLower(columnName), "'")
		}
		if columnName == "PRIMARY" {
			continue
		}

		if columnType == "text" {
			tagStr = CompactStr("text ", tagStr)
		}

		table.Columns = append(table.Columns, Column{
			Name:    columnName,
			Type:    SqlTypeToGoType(columnType),
			Tag:     tagStr,
			Comment: comment,
		})
	}

	return table, nil
}

func SqlTypeToGoType(sqlType string) string {
	switch strings.ToLower(sqlType) {
	case "int", "integer":
		return "int"
	case "bigint":
		return "int64"
	case "varchar", "text", "char":
		return "string"
	case "datetime", "timestamp":
		return "time.Time"
	case "float", "double":
		return "float64"
	default:
		return "string"
	}
}
