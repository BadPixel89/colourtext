package colourtext

import (
	"fmt"
	"os"
)

const Success = "[pass] "
const Fail = "[fail] "
const Error = "[err ] "
const Info = "[info] "
const Warn = "[warn] "
const Time = "[time] "
const Done = "[done] "
const Exit = "[exit] "

const Reset = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Magenta = "\033[35m"
const Cyan = "\033[36m"
const Gray = "\033[37m"
const White = "\033[97m"

func PrintlnWithColour(colour string, text string) {
	if !noColour() {
		fmt.Printf("%s%s%s\n", colour, text, Reset)
		return
	}
	fmt.Println(text)
}
func PrintColour(colour string, text string) {
	if !noColour() {
		fmt.Printf( "%s%s%s", colour, text, Reset)
		return
	}
	fmt.Print(text)
}

// prints a new line before the data, intention is for printing structs in colour
func PrintAny(colour string, v ...any) {
	fmt.Printf("%s%+v%s\n", colour, v, Reset)
}

// prepend [pass] print in green, prints new line
func PrintSuccess(text string) {
	PrintlnWithColour(Success+text, Green)
}

// prepend [fail] print in red, prints new line
func PrintFail(text string) {
	PrintlnWithColour(Fail+text, Red)
}

// prepend [err ] print in red, prints new line
func PrintError(text string) {
	PrintlnWithColour(Error+text, Red)
}

// prepend [info] print in cyan, prints new line
func PrintInfo(text string) {
	PrintlnWithColour(Info+text, Cyan)
}

// prepend [warn] print in yellow, prints new line
func PrintWarn(text string) {
	PrintlnWithColour(Warn+text, Yellow)
}

// prepend [time] print in cyan, prints new line
func PrintTime(text string) {
	PrintlnWithColour(Time+text, Cyan)
}

// prepend [done] print in green, prints new line
func PrintDone(text string) {
	PrintlnWithColour(Done+text, Green)
}

// prepend [exit] print in red, prints new line
func PrintExit(text string) {
	PrintlnWithColour(Exit+text, Red)
}

func noColour() bool {
	return os.Getenv("NO_COLOR") != "" && os.Getenv("NO_COLOR") != "0"
}
