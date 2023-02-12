package logger

import (
	"fmt"

	"github.com/go-niom/niom/pkg/constants"
	"github.com/gookit/color"
)

func Info(msg string) {
	color.Bluep(constants.AppSign)
	color.Greenp(" [INFO] ")
	fmt.Println(msg)
}

func Error(msg, err string) {
	color.Bluep(constants.AppSign)
	color.Redp(" [ERROR] ")
	fmt.Println(msg, err)
}

func Warn(msg string) {
	color.Bluep(constants.AppSign)
	color.Yellowp(" [WARN] ")
	fmt.Println(msg)
}
