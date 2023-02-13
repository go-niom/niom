package utils

import (
	"strings"
)

type ArgsResult struct {
	App  string
	Args []string
}

type ArgsStruct struct {
	Prefix string
	Args   []string
	Result ArgsResult
}

func ReadArgs(prefix string, args []string) string {
	for _, v := range args {
		if strings.HasPrefix(v, prefix) {
			return strings.TrimPrefix(v, prefix)
		}
	}
	return ""
}

func GetArgs(prefix string, args []string) []string {
	return strings.Split(ReadArgs(prefix, args), " ")
}

func (a *ArgsStruct) AppAndArgs() {
	cmd := GetArgs(a.Prefix, a.Args)
	if cmd != nil {
		a.Result.App = cmd[0]
		a.Result.Args = cmd[1:]
	}
}
