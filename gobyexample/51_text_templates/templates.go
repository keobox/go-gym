package main

import (
	"os"
	"text/template"
)

func main() {
	// A template object named t1
	t1 := template.New("t1")

	// Parse a value from a string
	// Basically assigns the string template to the template object
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// Same as above, but Must panics in case of error
	t1 = template.Must(t1.Parse("Value is {{.}}\n"))

	// By executing a template we generate test
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Python",
		"C",
	})

	// An helper
	Create := func(name, templ string) *template.Template {
		return template.Must(template.New("name").Parse(templ))
	}

	// In case of a struct is possible to access
	// fields using .Field
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// Templates works with maps too
	// Keys can be lowercase (but it's upper in the template string above,
	// to work is necessary to change the template string)
	t2.Execute(os.Stdout, map[string]string{"Name": "Mickey Mouse"})

	// The template string can have conditionals
	// something is false if his value is the default value
	// 0 for in, "" for string etc.
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// Is possible to have range blocks in the template string
	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"Go", "Python", "C"})
}
