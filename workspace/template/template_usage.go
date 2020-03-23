package main

import (
	"os"
	"text/template"
)

var templateContents = `##Template Test : {{.title}}
This is a test templates :(
## range .Members >>
{{range .Members}}Member > Name::{{.Name}} , Age :: {{.Age}}
{{end}}
## range .Members with index >>
{{range $index, $member := .Members}}[#{{$index}}] Member > Name::{{$member.Name}} , Age :: {{$member.Age}}
{{end}}
`

func main() {
	// given
	title := "Title for tests!"
	type member struct {
		Name string
		Age  int
	}
	var members []member
	members = append(members, member{"member1", 1})
	members = append(members, member{"member2", 2})
	members = append(members, member{"member3", 3})

	// binds
	bind := make(map[string]interface{})
	bind["title"] = title
	bind["Members"] = members

	tmpl, err := template.New("").Parse(templateContents)
	if err != nil {
		panic(err)
	}
	_ = tmpl.Execute(os.Stdout, bind)
}
