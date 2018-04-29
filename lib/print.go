package lib

import (
	"github.com/spiral/goffli/utils"
	"github.com/yuin/gopher-lua"
)

// Print replaces default function in lua with Printf analog plus `<red+hb>color formatting support</reset>`.
func Print(l *lua.LState) int {
	var (
		format string
		args   []interface{}
	)

	for i := 1; i <= l.GetTop(); i++ {
		if i == 1 {
			format = l.ToString(i)
			continue
		}

		args = append(args, l.ToString(i))
	}

	utils.Printf(format, args...)

	return 0
}
