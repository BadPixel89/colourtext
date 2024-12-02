## Basic coloured printing for console logging, with or without status text.

To use the colour functions on Windows in either Powershell or CMD (not Windows Terminal), you must first enable console virtualisation using the EnableVirtualConsoleStdout() and EnableVirtualConsoleStderr() functions. Otherwise the ANSI escape codes will be printed into the output as plain text. 

On Mac and Linux these functions do nothing so you don't need to check the OS in your code.
