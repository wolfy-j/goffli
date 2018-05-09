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

package lib

import (
	"errors"
	"fmt"
	"github.com/wolfy-j/goffli/utils"
	"github.com/yuin/gopher-lua"
	"gopkg.in/AlecAivazis/survey.v1"
	"os"
	"strconv"
)

// Prompter wraps set of default values (stack) and creates lua function to retrieve such values or request user input.
type Prompter struct {
	Args  []string
	stack int
}

// NewPrompter returns new lua function which works as prompt and associated with set of defaults values provided as args.
func NewPrompter(args []string) func(l *lua.LState) int {
	return (&Prompter{Args: args}).handler
}

func (p *Prompter) handler(l *lua.LState) int {
	label := l.ToString(1)
	defaultValue := l.ToString(3)

	if label == "" {
		panic("label is required")
	}

	if defaultValue != "" {
		label = fmt.Sprintf("%s [%s]", label, defaultValue)
	}

	// progress to next input position
	defer func() { p.stack++ }()

	validator := p.newValidator(l.ToString(2))

	var value string
	if len(p.Args) > p.stack {
		value = p.Args[p.stack]
		if err := validator(value); err != nil {
			value = "" // nope
		} else {
			utils.Printf("<green+hb>✔</reset> <white+hb>%s</reset> <cyan>%s</reset>\n", label, p.Args[p.stack])
		}
	}

	if value == "" {
		prompt := &survey.Input{Message: label}
		survey.AskOne(prompt, &value, validator)
	}

	if value == "" && defaultValue != "" {
		value = defaultValue
	}

	l.Push(lua.LString(value))
	return 1
}

// newValidator creates validator for various prompts (defaults to non empty string).
func (p *Prompter) newValidator(t string) func(ans interface{}) error {
	switch t {
	case "exists":
		return func(ans interface{}) error {
			if _, err := os.Stat(ans.(string)); err != nil {
				return errors.New("no such file")
			}

			return nil
		}

	case "int", "int64", "number":
		return func(ans interface{}) error {
			if _, err := strconv.ParseInt(ans.(string), 10, 64); err != nil {
				return errors.New("invalid number")
			}

			return nil
		}

	case "float", "float64":
		return func(ans interface{}) error {
			if _, err := strconv.ParseFloat(ans.(string), 64); err != nil {
				return errors.New("invalid number")
			}

			return nil
		}
	case "not-empty", "not_empty", "!empty":
		return func(ans interface{}) error {
			if ans.(string) == "" {
				return errors.New("empty value")
			}

			return nil
		}
	}

	return func(ans interface{}) error {
		return nil
	}
}
