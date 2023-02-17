package utils

import (
	"strings"
)

type ArgsResult struct {
	App  string
	Args []string
}

// ArgsResult holds niom start data
type ArgsStruct struct {
	Prefix string
	Args   []string
	Result ArgsResult
}

// ReadArgs reads arg from passed args as per the specified name
// for example `ReadArgs("-p=",[]string{"up","-p='db/migrations'","seed"})“
// will return `db/migrations`
func ReadArgs(prefix string, args []string) string {
	for _, v := range args {
		if strings.HasPrefix(v, prefix) {
			return strings.TrimPrefix(v, prefix)
		}
	}
	return ""
}

// GetArgs reads arg from passed args as per the specified name
// for example `ReadArgs("-c=",[]string{"up","-c='go run main.go'","seed"})“
// will return []string{"go", "run", "main.go"}
func GetArgs(prefix string, args []string) []string {
	return strings.Split(ReadArgs(prefix, args), " ")
}

// AppAndArgs sets ArgsResult
// for example `ArgsStruct{Prefix: "-c=",Args : []string{"up","-c='go run main.go'",}“
// will set ArgsResult{App: 'go' Args: []string{ "run", "main.go"}}
func (a *ArgsStruct) AppAndArgs() {
	cmd := GetArgs(a.Prefix, a.Args)
	if cmd != nil {
		a.Result.App = cmd[0]
		a.Result.Args = cmd[1:]
	}
}
