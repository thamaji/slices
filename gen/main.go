package main

import (
	"fmt"
	"os"
	"text/template"
	"unicode"
)

func main() {
	templ := template.New("xlice.tpl")

	templ.Funcs(template.FuncMap{
		"ToUpper": func(v string) string {
			return string(unicode.ToUpper(rune(v[0]))) + v[1:]
		},
	})

	if _, err := templ.ParseFiles("xlice.tpl"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	types := []string{
		"bool",
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64",
		"complex64", "complex128",
		"string",
		"byte",
		"rune",
	}

	for _, t := range types {
		f, err := os.Create("../" + t + ".go")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		ctx := map[string]interface{}{
			"Type": t,
		}

		err = templ.Execute(f, ctx)
		f.Close()

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
