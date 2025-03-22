package colourtext

import (
	"fmt"
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

func printlnWithColour(colour string, text string) {
	fmt.Printf("%s%s%s\n", colour, text, Reset)
}

// use colortext.Red as the colourstring for example
func PrintColor(colour string, text string) {
	printlnWithColour(text, colour)
}

// prints a new line before the data, intention is for printing structs in colour
func PrintAny(colour string, v ...any) {
	fmt.Printf("%s%+v%s\n", colour, v, Reset)
}

// prepend [pass] print in green
func PrintSuccess(text string) {
	printlnWithColour(Success+text, Green)
}

// prepend [fail] print in red
func PrintFail(text string) {
	printlnWithColour(Fail+text, Red)
}

// prepend [err ] print in red
func PrintError(text string) {
	printlnWithColour(Error+text, Red)
}

// prepend [info] print in cyan
func PrintInfo(text string) {
	printlnWithColour(Info+text, Cyan)
}

// prepend [warn] print in yellow
func PrintWarn(text string) {
	printlnWithColour(Warn+text, Yellow)
}

// prepend [time] print in cyan
func PrintTime(text string) {
	printlnWithColour(Time+text, Cyan)
}

// prepend [done] print in green
func PrintDone(text string) {
	printlnWithColour(Done+text, Green)
}

// prepend [exit] print in red
func PrintExit(text string) {
	printlnWithColour(Exit+text, Red)
}
