package logger

import (
	"fmt"
	"time"

	"github.com/go-niom/niom/pkg/constants"
	"github.com/gookit/color"
)

// It is the helper function to print current date-time in the logger
func printCurrentTime() {
	fmt.Print(time.Now().Format(" 2006-01-02 15:04:05"))
}

// Func Info log the msg passed to it
// With Indicator INFO in green color
func Info(msg string) {
	color.Bluep(constants.AppSign)
	printCurrentTime()
	color.Greenp(" [INFO] ")
	fmt.Println(msg)
}

// Func Error log the msg and error passed to it
// With Indicator Error in red color
func Error(msg, err string) {
	color.Bluep(constants.AppSign)
	printCurrentTime()
	color.Redp(" [ERROR] ")
	fmt.Println(msg, err)
}

// Func Warn log the msg passed to it
// With Indicator Warn in yellow color
func Warn(msg string) {
	color.Bluep(constants.AppSign)
	printCurrentTime()
	color.Yellowp(" [WARN] ")
	fmt.Println(msg)
}
