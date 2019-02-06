package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type param struct {
	Type      string
	RawType   string
	ZeroValue string
}

func main() {
	templ := template.New("opts.tpl")

	if _, err := templ.ParseFiles("opts.tpl"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	params := []param{
		param{Type: "Bool", RawType: "bool", ZeroValue: "false"},
		param{Type: "Int", RawType: "int", ZeroValue: "0"},
		param{Type: "Int8", RawType: "int8", ZeroValue: "0"},
		param{Type: "Int16", RawType: "int16", ZeroValue: "0"},
		param{Type: "Int32", RawType: "int32", ZeroValue: "0"},
		param{Type: "Int64", RawType: "int64", ZeroValue: "0"},
		param{Type: "Uint", RawType: "uint", ZeroValue: "0"},
		param{Type: "Uint8", RawType: "uint8", ZeroValue: "0"},
		param{Type: "Uint16", RawType: "uint16", ZeroValue: "0"},
		param{Type: "Uint32", RawType: "uint32", ZeroValue: "0"},
		param{Type: "Uint64", RawType: "uint64", ZeroValue: "0"},
		param{Type: "Uintptr", RawType: "uintptr", ZeroValue: "0"},
		param{Type: "Float32", RawType: "float32", ZeroValue: "0"},
		param{Type: "Float64", RawType: "float64", ZeroValue: "0"},
		param{Type: "Complex64", RawType: "complex64", ZeroValue: "0"},
		param{Type: "Complex128", RawType: "complex128", ZeroValue: "0"},
		param{Type: "String", RawType: "string", ZeroValue: "\"\""},
		param{Type: "Byte", RawType: "byte", ZeroValue: "0"},
		param{Type: "Rune", RawType: "rune", ZeroValue: "0"},
		param{Type: "Interface", RawType: "interface{}", ZeroValue: "nil"},
	}

	for _, param := range params {
		f, err := os.Create("../" + strings.ToLower(param.Type) + ".go")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		err = templ.Execute(f, param)
		f.Close()

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
