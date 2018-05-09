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
	"fmt"
	"github.com/pkg/errors"
	"github.com/wolfy-j/goffli/ffmpeg"
	"github.com/wolfy-j/goffli/utils"
	"github.com/yuin/gopher-lua"
	"time"
)

func ffmpegModule(L *lua.LState) int {
	L.Push(L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"probe":       probe,
		"run":         run,
		"showInfo":    nil,
		"showStreams": nil,
	}))

	return 1
}

func probe(l *lua.LState) int {
	utils.Log("probing", l.ToString(1))

	info, err := ffmpeg.Probe(l.ToString(1))
	if err != nil {
		panic(errors.WithMessage(err, "ffmpeg.probe"))
	}

	lv, err := Encode(info, l)
	if err != nil {
		panic(err)
	}

	// render media header
	if l.OptBool(2, false) {
		RenderMedia(info)
		fmt.Print("\n")
	}

	// render all streams
	if l.OptBool(3, false) {
		RenderStreams(info)
		fmt.Print("\n")
	}

	l.Push(lv)

	return 1
}

func run(l *lua.LState) int {
	opts := make([]string, 0)
	if err := Decode(l.Get(1), &opts); err != nil {
		panic(err)
	}

	utils.Log("ffmpeg", opts...)

	var pb *utils.Progress

	err := ffmpeg.Run(opts, func(c, t time.Duration) {
		if pb == nil {
			pb = utils.NewProgress(t, l.ToString(2))
		}

		pb.Set(c)
	})

	if pb != nil {
		pb.Finish()
	}

	if err != nil {
		panic(errors.WithMessage(err, "ffmpeg.run"))
	}

	return 0
}
