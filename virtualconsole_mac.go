//go:build darwin

package colourtext

// enables console virtualisation in Windows, allowing PowerShell and CMD to output coloured text.-
// if this is not enabled, any output will have the colour characters prepended and appended to it
func EnableVirtualConsoleStdout() error {
	return nil
}

func EnableVirtualConsoleStderr() error {
	return nil
}
