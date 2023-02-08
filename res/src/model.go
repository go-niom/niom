package src

var Model = `package user

type {{ .Name}} struct {
	Sample string` + "`json:\"sample,omitempty\"`\n" +
	`}
`
