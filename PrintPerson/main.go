package main

import (
	"html/template"
	"os"
	"log"
)

type Person struct {
	Name string
	Age int
	Emails []string
	Jobs []*Job
}

type Job struct {
	Employer string
	Role string
}

const templ  = `The name is {{.Name}}.
The age is {{.Age}}.
{{range .Emails}}
	An email is {{.}}
{{end}}
{{with .Jobs}}
	{{range .}}
		An employer is {{.Employer}}
		and the role is {{.Role}}
	{{end}}
{{end}}
`

func main() {
	job1 := Job{Employer: "FPT", Role: "Student"}
	job2 := Job{Employer: "TCW", Role: "Software Enginer"}
	
	person := Person{
		Name: "Danh Tran",
		Age: 23,
		Emails: []string{"danh@gmail.com", "abc@gmail.com"},
		Jobs: []*Job{&job1, &job2},
	}
	
	t := template.New("Person Template")
	t, err := t.Parse(templ)
	checkError(err)
	
	err = t.Execute(os.Stdout, person)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Error: ", err.Error())
		os.Exit(1)
	}
}
