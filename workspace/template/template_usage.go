package main

import (
	"text/template"
)

var templateContents = ` Template Test : {{.Title}}
	This is a test templates :(
`

func main() {
	bind := make(map[string]string)
	bind["Title"] = "test"

	tmpl, err := template.New("").Parse(templateContents)
	if err != nil {
		panic(err)
	}
}
