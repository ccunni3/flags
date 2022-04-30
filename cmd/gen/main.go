package main

import (
	"flag"
	"os"
	"text/template"
)

//go:generate sh -c "go run main.go -name=String -type=string > ../../strings.go"
type data struct {
	T     string
	TName string
}

func main() {
	var d data

	flag.StringVar(&d.T, "type", "", "The type.")
	flag.StringVar(&d.TName, "name", "", "The name of the type.")

	flag.Parse()

	t := template.Must(template.New("queue").Parse(tpl))
	t.Execute(os.Stdout, d)
}

var tpl = `// This file is autogenerated by cmd/gen DO NOT EDIT
package flags

// {{.TName}}Value flag type.
type {{.TName}}Value struct {
	Value
	V *{{.T}}
}

// {{.TName}} creates new {{.TName}} flag.
// Accepts a list of additional resolvers that are evaluated in sequence and
// the first one to yield a valid value is chosen.
// If no resolver yileds a valid value the default flag value is used.
// If flag is provided as a cli arg it will take precedance over all resolvers and the default value.
func (fs *FlagSet) {{.TName}}(name, usage {{.T}}, val {{.T}}, r ...ResolverFunc) *{{.T}} {
	fs.initFlagSet()

	v := {{.TName}}Value{
		Value: Value{
			name:      name,
			resolvers: r,
		},
		V: fs.fs.{{.TName}}(name, val, usage),
	}

	fs.Values = append(fs.Values, v)

	return v.V
}

func (fs *FlagSet) parse{{.TName}}Vals() {
	for i, val := range fs.Values {
		{{.T}}Val, ok := val.({{.TName}}Value)
		if !ok {
			continue
		}

		if fs.hasArg({{.T}}Val.name) {
			continue
		}

		for _, r := range {{.T}}Val.resolvers {
			if r(fs, {{.T}}Val.name, "", i) {
				break
			}
		}
	}
}
`