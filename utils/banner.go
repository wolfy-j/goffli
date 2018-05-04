package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// Banner allows automatic line alignment within desired width, banner includes color rendering.
type Banner struct {
	width int
	lines []string
}

// NewBanner returns new banner of desired width.
func NewBanner(width int) *Banner {
	return &Banner{width: width, lines: make([]string, 0)}
}

// Add adds new line to banner
func (b *Banner) Add(format string, args ...interface{}) {
	b.lines = append(b.lines, fmt.Sprintf(format, args...))
}

// Render returns compiled banner string.
func (b *Banner) String() string {

	var result string
	for _, l := range b.lines {
		ln := clearLen(l)

		if ln < b.width {
			result += fmt.Sprintf("%s%s\n", strings.Repeat(" ", (b.width-ln)/2), l)
		} else {
			result += fmt.Sprintf("%s\n", l)
		}
	}

	return strings.Trim(Sprintf(result), "\n")
}

// returns line width without any coloring tags
func clearLen(line string) int {
	r, err := regexp.Compile(`<([^>]+)>`)

	if err != nil {
		panic(err)
	}

	return len(r.ReplaceAllString(line, ""))
}
