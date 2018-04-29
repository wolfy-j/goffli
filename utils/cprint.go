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

package utils

import (
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1/core"
	"regexp"
	"strings"
)

// Printf works identically to fmt.Print but adds `<white+hb>color formatting support for CLI</reset>`.
func Printf(format string, args ...interface{}) {
	fmt.Print(Sprintf(format, args...))
}

// Sprintf works identically to fmt.Sprintf but adds `<white+hb>color formatting support for CLI</reset>`.
func Sprintf(format string, args ...interface{}) string {
	r, err := regexp.Compile(`<([^>]+)>`)
	if err != nil {
		panic(err)
	}

	format = r.ReplaceAllStringFunc(format, func(s string) string {
		return fmt.Sprintf(`{{color "%s"}}`, strings.Trim(s, "<>/"))
	})

	out, err := core.RunTemplate(fmt.Sprintf(format, args...), nil)
	if err != nil {
		panic(err)
	}

	return out
}
