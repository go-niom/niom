package main

import (
	"fmt"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

// list questions to ask
var qs = []*survey.Question{
	{
		Name:      "name",
		Prompt:    &survey.Input{Message: "Enter your name?"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "pet",
		Prompt: &survey.MultiSelect{
			Message: "Choose your pets:",
			Options: []string{"dogs", "reptiles", "cats", "birds", "fish", "rabbits", "pigs", "rats", "mices"},
			Default: "dogs",
		},
	},
	{
		Name:   "rating",
		Prompt: &survey.Input{Message: "Rate our website (integer number):"},
	},
}

func main2() {
	// store answer to struct
	// answer := struct {
	// 	Name   string   // survey match the question and field names
	// 	Pet    []string `survey:"pet"` //tag fields to match a specific name
	// 	Rating int      // if the types don't match, survey will convert it
	// }{}

	// err := survey.Ask(qs, &answer)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// petString := strings.Join(answer.Pet, ", ")

	// fmt.Printf("%s likes %s.", answer.Name, petString)

	// cmdExecute("cd", []string{"data"})
	// cmdExecute("go", []string{"mod", "tidy"})
	// cmdExecute("go", []string{"run", "main.go"})

}

func cmdExecute(dir, app string, args []string) {
	println(app)
	var cmd *exec.Cmd
	if len(args) > 1 {
		cmd = exec.Command(app, args...)
	} else {
		cmd = exec.Command(app)
	}
	cmd.Dir = dir
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}
