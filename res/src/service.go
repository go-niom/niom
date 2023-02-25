package src

var ServiceTmpl = `package {{ .NameLowerCase}}

import (
	"fmt"

	"{{ .ModuleName}}/pkg/logger"
)

var {{ .Name}}Service {{ .NameLowerCase}}ServiceInterface = &{{ .NameLowerCase}}Service{}

type {{ .NameLowerCase}}Service struct {
}

type {{ .NameLowerCase}}ServiceInterface interface {
	create(interface{}) (string, interface{})
	findAll(interface{}) (string, interface{})
	findOne(string) (string, interface{})
	update(string, interface{}) (string, interface{})
	remove(string) (string, interface{})
}

func (a *{{ .NameLowerCase}}Service) create(data interface{}) (string, interface{}) {
	logger.Info("Service : {{ .Name}}, Method : create")
	return "This action adds a new {{ .NameLowerCase}}", nil
}

func (a *{{ .NameLowerCase}}Service) findAll(data interface{}) (string, interface{}) {
	logger.Info("Service : {{ .Name}}, Method : findAll")
	return "This action find all {{ .NameLowerCase}}", nil
}

func (a *{{ .NameLowerCase}}Service) findOne(id string) (string, interface{}) {
	logger.Info("Service : {{ .Name}}, Method : findOne")
	return fmt.Sprintf("This action returns a #%s {{ .NameLowerCase}}", id), nil
}

func (a *{{ .NameLowerCase}}Service) update(id string, data interface{}) (string, interface{}) {
	logger.Info("Service : {{ .Name}}, Method : update")
	return fmt.Sprintf("This action updates a #%s user", id), nil
}

func (a *{{ .NameLowerCase}}Service) remove(id string) (string, interface{}) {
	logger.Info("Service : {{ .Name}}, Method : remove")
	return fmt.Sprintf("This action removes a #%s user", id), nil
}

`

var ServiceEmptyTmpl = `package {{ .PackageName}}

import (
	"fmt"

	"{{ .ModuleName}}/pkg/logger"
)

var {{ .Name}}Service {{ .NameLowerCase}}ServiceInterface = &{{ .NameLowerCase}}Service{}

type {{ .NameLowerCase}}Service struct {
}

type {{ .NameLowerCase}}ServiceInterface interface {
	create(interface{}) (string, interface{})
	findAll(interface{}) (string, interface{})
	findOne(string) (string, interface{})
	update(string, interface{}) (string, interface{})
	remove(string) (string, interface{})
}

func (a *{{ .NameLowerCase}}Service) create(data interface{}) (string, interface{}) {
	logger.Info("Service : {{ .Name}}, Method : create")
	return "This action adds a new {{ .NameLowerCase}}", nil
}

func (a *{{ .NameLowerCase}}Service) findAll(data interface{}) (string, interface{}) {
	logger.Info("Service : {{ .Name}}, Method : findAll")
	return "This action find all {{ .NameLowerCase}}", nil
}

func (a *{{ .NameLowerCase}}Service) findOne(id string) (string, interface{}) {
	logger.Info("Service : {{ .Name}}, Method : findOne")
	return fmt.Sprintf("This action returns a #%s {{ .NameLowerCase}}", id), nil
}

func (a *{{ .NameLowerCase}}Service) update(id string, data interface{}) (string, interface{}) {
	logger.Info("Service : {{ .Name}}, Method : update")
	return fmt.Sprintf("This action updates a #%s user", id), nil
}

func (a *{{ .NameLowerCase}}Service) remove(id string) (string, interface{}) {
	logger.Info("Service : {{ .Name}}, Method : remove")
	return fmt.Sprintf("This action removes a #%s user", id), nil
}

`
