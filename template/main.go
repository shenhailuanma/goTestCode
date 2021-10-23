package main

import (
	"bytes"
	"fmt"
	"text/template"
)

const templateText1 = `
=== template start ===
{{- if exist .BoolTrue}}
exist = {{exist .BoolTrue}}
BoolTrue = {{.BoolTrue}}
{{- end}}
{{- if exist .BoolFalse}}
exist = {{exist .BoolFalse}}
BoolFalse = {{.BoolFalse}}
{{- end}}
{{- if exist .NotExist}}
NotExist = {{.NotExist}}
{{- end}}
existNotExist = {{exist .NotExist}}
=== template end ===
`
var templateData1 = map[string]interface{}{
	"BoolTrue" : true,
	"BoolFalse" : false,
}

func main() {
	fmt.Println("start")

	output, err := templateRender(templateText1, templateData1)
	if err != nil {
		fmt.Println("templateRender error:", err.Error())
	}
	fmt.Println("templateText1, output:")
	fmt.Println(output)


	fmt.Println("end")
}

func keyExist(input interface{}) bool {
	fmt.Println(input)
	if input != nil {
		return true
	}
	return false
}

func templateRender(templateText string, inputs map[string]interface{}) (string, error) {

	tmpl := template.New("test").Funcs(template.FuncMap{
		"exist": keyExist,
	})
	tmpl, err := tmpl.Parse(templateText)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, inputs)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
