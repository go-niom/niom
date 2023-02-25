package src

var DTO = `package dto

type Create{{ .Name}}Dto struct {
	Sample string` + "`json:\"sample,omitempty\"`\n" +
	`}

type Update{{ .Name}}Dto struct {
	Sample string` + "`json:\"sample,omitempty\"`\n" +
	`}

type Query{{ .Name}}Dto struct {
	Sample string` + "`json:\"sample,omitempty\"`\n" +
	`}

`
