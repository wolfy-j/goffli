package utils

import "strings"

// Verbose enabled or disabled showing debug messages in CLI.
var Verbose bool

// Log adds debug message to the output.
func Log(name string, args ...string) {
	if !Verbose {
		return
	}

	Printf("<cyan+hb>►</reset> <yellow+hb>%s</reset> <green+hb>%s</reset>\n", name, strings.Join(args, " "))
}
