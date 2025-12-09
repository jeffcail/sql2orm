package handler

import (
	"encoding/json"
	"go/format"
	"io"
	"net/http"
	"strings"

	"github.com/jeffcail/sql2orm/templates"
	"github.com/jeffcail/sql2orm/utils"
)

type RequestBody struct {
	SQL string `json:"sql"`
	Typ int    `json:"typ"`
}

type Response struct {
	Struct string `json:"struct"`
}

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	bodyBytes, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	req := new(RequestBody)
	if err := json.Unmarshal(bodyBytes, req); err != nil {
		jsonError(w, err.Error())
		return
	}

	table, err := utils.ParseSQLTable(req.SQL)
	if err != nil {
		jsonError(w, err.Error())
		return
	}

	var tmplSet *_templates.TemplateSet
	if req.Typ == 1 {
		tmplSet, err = _templates.ParseXORM()
	} else {
		tmplSet, err = _templates.ParseGORM()
	}
	if err != nil {
		jsonError(w, "模板解析失败: "+err.Error())
		return
	}

	data := struct {
		Tables []utils.Table
	}{Tables: []utils.Table{table}}

	var code strings.Builder

	code.WriteString("package models\n\n")

	if tmplSet.ImportTime {
		code.WriteString("import (\n\t\"time\"\n)\n\n")
	}

	if err = tmplSet.Struct.Execute(&code, data); err != nil {
		jsonError(w, "生成 struct 失败: "+err.Error())
		return
	}

	code.WriteString("\n")
	if err = tmplSet.TableName.Execute(&code, data); err != nil {
		jsonError(w, "生成 TableName 方法失败: "+err.Error())
		return
	}

	result := code.String()
	if idx := strings.LastIndex(result, "ENGINE"); idx != -1 {
		if start := strings.LastIndex(result[:idx], "\n"); start != -1 {
			result = result[:start+1] // +1 保留换行
		}
	}

	formatted, err := format.Source([]byte(result))
	if err != nil {
		jsonError(w, "Go 代码格式化失败:\n"+err.Error()+"\n\n原始代码:\n"+result)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp := Response{Struct: string(formatted)}
	json.NewEncoder(w).Encode(resp)
}

func jsonError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	resp := Response{Struct: "错误: " + msg}
	_ = json.NewEncoder(w).Encode(resp)
}
