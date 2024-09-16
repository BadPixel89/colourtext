package colourtext

import (
	"log"
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

func loglnWithColour(colour string, text string) {
	log.Printf("%s%s%s\n", colour, text, Reset)
}
func logWithColour(colour string, text string) {
	log.Printf("%s%s%s", colour, text, Reset)
}

// use colortext.Red as the colourstring for example
func PrintColour(colour string, text string) {
	logWithColour(colour, text)
}

// use colortext.Red as the colourstring for example
func PrintlnColour(colour string, text string) {
	loglnWithColour(text, colour)
}

// prints a new line before the data, intention is for printing structs in colour
func PrintAny(colour string, v ...any) {
	log.Printf("%s%+v%s", colour, v, Reset)
}

// prepend [pass] print in green
func PrintSuccess(text string) {
	loglnWithColour(Green, Success+text)
}

// prepend [fail] print in red
func PrintFail(text string) {
	loglnWithColour(Red, Fail+text)
}

// prepend [err ] print in red
func PrintError(text string) {
	loglnWithColour(Red, Error+text)
}

// prepend [info] print in cyan
func PrintInfo(text string) {
	loglnWithColour(Cyan, Info+text)
}

// prepend [warn] print in yellow
func PrintWarn(text string) {
	loglnWithColour(Yellow, Warn+text)
}

// prepend [time] print in cyan
func PrintTime(text string) {
	loglnWithColour(Cyan, Time+text)
}

// prepend [done] print in green
func PrintDone(text string) {
	loglnWithColour(Green, Done+text)
}

// prepend [exit] print in red
func PrintExit(text string) {
	loglnWithColour(Red, Exit+text)
}
