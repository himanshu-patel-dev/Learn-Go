// Go offers built-in support for creating dynamic content or
// showing customized output to the user with the text/template
// package. A sibling package named html/template provides the
// same API but has additional security features and should be
// used for generating HTML.

package main

import (
	"html/template"
	"os"
)

func main() {
	// We can create a new template and parse its body from a string.
	// Templates are a mix of static text and “actions” enclosed in
	// {{...}} that are used to dynamically insert content.
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// By “executing” the template we generate its text with specific
	// values for its actions. The {{.}} action is replaced by the value
	// passed as a parameter to Execute.
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})
	// Value is some text
	// Value is 5
	// Value is [Go Rust C&#43;&#43; C#]

	// Alternatively, we can use the template.Must function to
	// panic in case Parse returns an error.
	// t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// method to return a template executor if we pass name and template
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// creating a new template with Create func
	t2 := Create("t2", "Name: {{.Name}}\n")

	// If the data is a struct we can use the {{.FieldName}} action to access
	// its fields. The fields should be exported to be accessible when a
	// template is executing.
	t2.Execute(os.Stdout, struct{ Name string }{"Jane Doe"})
	// Name: Jane Doe

	// The same applies to maps; with maps there is no restriction on
	// the case of key names.
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})
	// Name: Mickey Mouse

	// A value is considered false if it’s the default value of a type,
	// such as 0, an empty string, nil pointer, etc. This sample
	// demonstrates another feature of templates: using - in actions to
	// trim whitespace.
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty") // yes
	t3.Execute(os.Stdout, "")          // no

	// range blocks let us loop through slices, arrays, maps or channels.
	// Inside the range block {{.}} is set to the current item of the
	// iteration.

	t4 := Create("t4", "Range: \n{{range .}} - {{.}}\n{{end}}\n")
	t4.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})
	// Range:
	// 	- Go
	// 	- Rust
	// 	- C&#43;&#43;
	// 	- C#
}
