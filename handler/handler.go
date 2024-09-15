package handler

import (
	"encoding/json"
	"github.com/jeffcail/sql2orm/utils"
	"go/format"
	"io"
	"net/http"
	"strings"
	"text/template"
)

type RequestBody struct {
	SQL string `json:"sql"`
	Typ int    `json:"typ"`
}

type Response struct {
	Struct string `json:"struct"`
}

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有域
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	body := r.Body
	bytes, _ := io.ReadAll(body)

	rb := new(RequestBody)
	err := json.Unmarshal(bytes, rb)
	if err != nil {
		res := new(Response)
		res.Struct = err.Error()
		marshal, _ := json.Marshal(res)
		_, _ = w.Write(marshal)
	}

	table, err := utils.ParseSQLTable(rb.SQL)
	if err != nil {
		res := new(Response)
		res.Struct = err.Error()
		marshal, _ := json.Marshal(res)
		_, _ = w.Write(marshal)
		return
	}

	var structTmpl *template.Template
	var tableNameTmpl *template.Template
	if rb.Typ == 1 {
		structTmpl, tableNameTmpl, err = parseXorm()
		if err != nil {
			res := new(Response)
			res.Struct = err.Error()
			marshal, _ := json.Marshal(res)
			_, _ = w.Write(marshal)
			return
		}
	} else {
		structTmpl, tableNameTmpl, err = parseGorm()
		if err != nil {
			res := new(Response)
			res.Struct = err.Error()
			marshal, _ := json.Marshal(res)
			_, _ = w.Write(marshal)
			return
		}
	}

	var structStr strings.Builder
	err = structTmpl.Execute(&structStr, struct {
		Tables []utils.Table
	}{
		Tables: []utils.Table{table},
	})
	if err != nil {
		res := new(Response)
		res.Struct = err.Error()
		marshal, _ := json.Marshal(res)
		_, _ = w.Write(marshal)
		return
	}

	var tableNameStr strings.Builder
	err = tableNameTmpl.Execute(&tableNameStr, struct {
		Tables []utils.Table
	}{
		Tables: []utils.Table{table},
	})
	if err != nil {
		res := new(Response)
		res.Struct = err.Error()
		marshal, _ := json.Marshal(res)
		_, _ = w.Write(marshal)
		return
	}

	// Combine both parts
	s := structStr.String()
	contains := strings.Contains(s, "ENGINE")
	if contains {
		index := strings.LastIndex(s, "ENGINE")
		s = s[:index]
	}

	combinedStr := s + tableNameStr.String()

	// Format the combined code
	formattedStruct, err := format.Source([]byte(combinedStr))
	if err != nil {
		res := new(Response)
		res.Struct = err.Error()
		marshal, _ := json.Marshal(res)
		_, _ = w.Write(marshal)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	res := new(Response)
	res.Struct = string(formattedStruct)
	marshal, _ := json.Marshal(res)
	_, _ = w.Write(marshal)
	return
}
