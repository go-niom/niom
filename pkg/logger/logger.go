package logger

import (
	"fmt"
	"time"

	"github.com/go-niom/niom/pkg/constants"
	"github.com/gookit/color"
)

func printCurrentTime() {
	fmt.Print(time.Now().Format(" 2006-01-02 15:04:05"))
}
func Info(msg string) {
	color.Bluep(constants.AppSign)
	printCurrentTime()
	color.Greenp(" [INFO] ")
	fmt.Println(msg)
}

func Error(msg, err string) {
	color.Bluep(constants.AppSign)
	printCurrentTime()
	color.Redp(" [ERROR] ")
	fmt.Println(msg, err)
}

func Warn(msg string) {
	color.Bluep(constants.AppSign)
	printCurrentTime()
	color.Yellowp(" [WARN] ")
	fmt.Println(msg)
}
