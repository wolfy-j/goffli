package lib

import (
	"fmt"
	"git.spiralscout.com/wolfy-j/goffli/ffmpeg"
	"git.spiralscout.com/wolfy-j/goffli/utils"
	"github.com/pkg/errors"
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
	defer pb.Finish()

	err := ffmpeg.Run(opts, func(c, t time.Duration) {
		if pb == nil {
			pb = utils.NewProgress(t, l.ToString(2) == "spinner")
		}

		pb.Set(c)
	})

	if err != nil {
		panic(errors.WithMessage(err, "ffmpeg.run"))
	}

	return 0
}
