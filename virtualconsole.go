//go:build windows

package colourtext

import (
	"runtime"

	"golang.org/x/sys/windows"
)

// enables console virtualisation in Windows, allowing PowerShell and CMD to output coloured text.-
// if this is not enabled, any output will have the colour characters prepended and appended to it
func EnableVirtualConsole(console windows.Handle) error {
	if runtime.GOOS == "linux" {
		return nil
	}
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
