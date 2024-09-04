package colourtext

import (
	"log"

	"golang.org/x/sys/windows"
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

const _PARAMETER_IS_INCORRECT = 87

func EnableVirtualConsole(console windows.Handle) error {
	var mode uint32
	err := windows.GetConsoleMode(console, &mode)
	if err != nil {
		return err
	}
	err = windows.SetConsoleMode(console, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	if err != nil {
		return err
	}
	return nil
}

func loglnWithColour(text string, colour string) {
	log.Printf("%s%s%s\n", colour, text, Reset)
}

// use colortext.Red as the colourstring for example
func PrintColor(text string, colour string) {
	loglnWithColour(text, colour)
}

// prints a new line before the data, intention is for printing structs in colour
func PrintAny(colour string, v ...any) {
	log.Printf(colour+"\n%+v"+Reset, v...)
}

func PrintSuccess(text string) {
	loglnWithColour(Success+text, Green)
}
func PrintFail(text string) {
	loglnWithColour(Fail+text, Red)
}
func PrintError(text string) {
	loglnWithColour(Error+text, Red)
}
func PrintInfo(text string) {
	loglnWithColour(Info+text, Cyan)
}
func PrintWarn(text string) {
	loglnWithColour(Warn+text, Yellow)
}
func PrintTime(text string) {
	loglnWithColour(Time+text, Cyan)
}
func PrintDone(text string) {
	loglnWithColour(Done+text, Green)
}
func PrintExit(text string) {
	loglnWithColour(Exit+text, Red)
}
