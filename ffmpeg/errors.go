package ffmpeg

import (
	"fmt"
	"strings"
)

// last N lines from ffmpeg to count as error(s).
const errorLines = 1

// extractError trims ffmpeg output to last n lines only.
func extractError(err error, cmdErr string) error {
	if cmdErr == "" {
		return fmt.Errorf("ffmpeg error: %s", err)
	}

	lines := rmeLines(strings.Split(cmdErr, "\n"))
	return fmt.Errorf(
		"ffmpeg error (%s): %s",
		err,
		strings.Join(lines[len(lines)-errorLines:], "; "),
	)
}

// rmeLines removes all blank lines and trims all strings.
func rmeLines(lines []string) []string {
	result := make([]string, 0)

	for _, l := range lines {
		l = strings.Trim(l, " \n\r")
		if l != "" {
			result = append(result, l)
		}
	}

	return result
}
