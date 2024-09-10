//go:build windows

package colourtext

import (
	"runtime"

	"golang.org/x/sys/windows"
)

// enables console virtualisation in Windows, allowing PowerShell and CMD to output coloured text.-
// if this is not enabled, any output will have the colour characters prepended and appended to it
func EnableVirtualConsoleStdout() error {
	var mode uint32
	err := windows.GetConsoleMode(windows.Stdout, &mode)
	if err != nil {
		return err
	}
	err = windows.SetConsoleMode(windows.Stdout, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	if err != nil {
		return err
	}
	return nil
}

func EnableVirtualConsoleStderr() error {
	if runtime.GOOS == "linux" {
		return nil
	}
	var mode uint32
	err := windows.GetConsoleMode(windows.Stderr, &mode)
	if err != nil {
		return err
	}
	err = windows.SetConsoleMode(windows.Stderr, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	if err != nil {
		return err
	}
	return nil
}
