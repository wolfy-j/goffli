// Copyright © 2018 Wolfy-J <wolfy.jd@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ffmpeg

import (
	"fmt"
	"strings"
)

// last N lines from ffmpeg to count as error(s).
const errorLines = 3

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
