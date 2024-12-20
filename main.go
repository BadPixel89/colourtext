package colourtext

import (
	"log"
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

func loglnWithColour(colour string, text string) {
	if !noColour() {
		log.Printf("%s%s%s\n", colour, text, Reset)
		return
	}
	log.Printf("%s\n", text)
}
func logWithColour(colour string, text string) {
	if !noColour() {
		log.Printf("%s%s%s", colour, text, Reset)
		return
	}
	log.Printf("%s", text)
}

// use colortext.Red as the colourstring for example
func PrintColour(colour string, text string) {
	logWithColour(colour, text)
}

// use colortext.Red as the colourstring for example
func PrintlnColour(colour string, text string) {
	loglnWithColour(colour, text)
}

// prints a new line before the data, intention is for printing structs in colour
func PrintAny(colour string, v ...any) {
	log.Printf("%s%+v%s", colour, v, Reset)
}

// prepend [pass] print in green ends with newline
func PrintSuccess(text string) {
	loglnWithColour(Green, Success+text)
}

// prepend [fail] print in red ends with newline
func PrintFail(text string) {
	loglnWithColour(Red, Fail+text)
}

// prepend [err ] print in red ends with newline
func PrintError(text string) {
	loglnWithColour(Red, Error+text)
}

// prepend [info] print in cyan ends with newline
func PrintInfo(text string) {
	loglnWithColour(Cyan, Info+text)
}

// prepend [warn] print in yellow ends with newline
func PrintWarn(text string) {
	loglnWithColour(Yellow, Warn+text)
}

// prepend [time] print in cyan ends with newline
func PrintTime(text string) {
	loglnWithColour(Cyan, Time+text)
}

// prepend [done] print in green ends with newline
func PrintDone(text string) {
	loglnWithColour(Green, Done+text)
}

// prepend [exit] print in red ends with newline
func PrintExit(text string) {
	loglnWithColour(Red, Exit+text)
}

func noColour() bool {
	return os.Getenv("NO_COLOR") != ""
}
