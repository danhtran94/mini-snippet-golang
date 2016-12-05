package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Person struct {
	Name string
	Emails []string
}

const templ = `{{ $name := .Name -}}
The name is {{ $name }}.
{{ range .Emails -}}
	{{ $name }} has an emails is "{{ . | emailExpand }}"
{{ end }}`

func EmailExpander(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	return (substrs[0] +" at "+ substrs[1])
}

func main() {
	person := Person{
		Name: "Danh",
		Emails: []string{"danh@gmail.com", "dung@gmail.com"},
	}
	
	t := template.New("Person template")
	t = t.Funcs(template.FuncMap{"emailExpand": EmailExpander})
	
	t, err := t.Parse(templ)
	
	checkError(err)
	
	err = t.Execute(os.Stdout, person)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}

