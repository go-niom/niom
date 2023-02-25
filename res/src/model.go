package src

var Model = `package model

type {{ .Name}} struct {
	Sample string` + "`json:\"sample,omitempty\"`\n" +
	`}
`
