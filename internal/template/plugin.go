package template

var (
	Plugin = `package main
{{if .Plugins}}
import ({{range .Plugins}}
	_ "github.com/Augustu/go-plugins/{{.}}"{{end}}
){{end}}
`
)
