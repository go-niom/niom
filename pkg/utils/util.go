package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"os"
	"strings"
	"text/template"

	"github.com/gookit/color"
)

type TemplateArgs struct {
	Name          string
	NameLowerCase string
	ModuleName    string
	PackageName   string
}

// CreateFileWithData creates file at given directory with passed data
func CreateFileWithData(filePath, data string) {
	splitDir := strings.Split(filePath, "/")
	if len(splitDir) > 1 {
		dir := strings.Join(splitDir[0:len(splitDir)-1], "/")
		if err := ensureDir(dir); err != nil {
			fmt.Println("Directory creation failed with error: " + err.Error())
			os.Exit(1)
		}
	}
	file_name := strings.Join(splitDir, "/")
	file, err := os.Create(file_name)
	if err != nil {
		panic(err)
	}
	file.Write([]byte(data))
	file.Name()
	color.Println(`<green>CREATE</> file:`, file.Name())
	defer file.Close()
}

// GetAppName name from the moduleName
func GetAppName(moduleName string) string {
	split := strings.Split(moduleName, "/")
	return split[len(split)-1]
}

// ListContains check list has given value or not
func ListContains(value string, list []string) bool {
	for _, v := range list {
		if v == strings.Trim(value, " ") {
			return true
		}
	}
	return false
}

// UserPrompt accepts user input from the cli
func UserPrompt(message string) string {
	// Create a new scanner to read input from the command line
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(message)
	// Scan for the next token, which is assumed to be a line
	scanner.Scan()
	// Get the text from the scanned line
	input := scanner.Text()
	return strings.Trim(input, " ")
}

// RenderWriteToFile creates src files
func RenderWriteToFile(tmpl string, func_name string, file_name string) {
	td := TemplateArgs{Name: strings.Title(func_name), NameLowerCase: func_name}
	renderWrite(td, tmpl, file_name)
}

// RenderWriteToFileModule creates src module files
func RenderWriteToFileModule(tmpl, fileName, resName, module_name, packageName string) {
	td := TemplateArgs{Name: strings.Title(resName), NameLowerCase: resName, ModuleName: module_name, PackageName: packageName}
	renderWrite(td, tmpl, fileName)
}

// RenderMain creates main.go
func RenderMain(tmpl, fileName, packageName, module_name string) {
	td := TemplateArgs{Name: strings.Title(strings.Split(fileName, "/")[0]), NameLowerCase: packageName, ModuleName: module_name}
	renderWrite(td, tmpl, fileName)
}

// renderWrite renders the content and generate files at give directory
func renderWrite(td TemplateArgs, tmpl, file_name string) {
	splitDir := strings.Split(file_name, "/")
	if len(splitDir) > 1 {
		dir := strings.Join(splitDir[0:len(splitDir)-1], "/")
		if err := ensureDir(dir); err != nil {
			fmt.Println("Directory creation failed with error: " + err.Error())
			os.Exit(1)
		}
	}

	file_name = strings.Join(splitDir, "/")

	t, err := template.New("name").Parse(tmpl)
	if err != nil {
		fmt.Println("errrror", err)
	}
	f, err := os.Create(file_name)
	if err != nil {
		fmt.Println("create file: ", err)
		return
	}
	err = t.Execute(f, td)
	if err != nil {
		panic(err)
	}

	//Notify the user of the created files
	color.Println(`<green>CREATE</> file:`, file_name)
}

// Will be used in the future
func ReadAndCreateFile(source, appName, moduleName string) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	println(source)
	file, err := os.Open(exPath + "/" + source)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadAll(file)
	fmt.Print(b)
	// td := TemplateArgs{Name: strings.Title(packageName), NameLowerCase: packageName, ModuleName: module_name}
	// renderWrite(td, tmpl, fileName)
}

// ensureDir checks the passed path is exists or not
// if it is not exists then it creates path
func ensureDir(dirName string) error {

	err := os.MkdirAll(dirName, os.ModePerm)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}
