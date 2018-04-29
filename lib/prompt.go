package lib

import (
	"errors"
	"github.com/spiral/goffli/utils"
	"github.com/yuin/gopher-lua"
	"gopkg.in/AlecAivazis/survey.v1"
	"os"
	"strconv"
	"strings"
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
	label := strings.Title(l.ToString(1))

	if label == "" {
		panic("label is required")
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
	}

	return func(ans interface{}) error {
		if ans.(string) == "" {
			return errors.New("empty value")
		}

		return nil
	}
}
