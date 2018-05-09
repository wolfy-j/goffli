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

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wolfy-j/goffli/ffmpeg"
	"github.com/wolfy-j/goffli/lib"
	"github.com/wolfy-j/goffli/utils"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "info [path]",
		Short: "Show media file information",
		Args:  cobra.ExactArgs(1),
		Run:   infoHandler,
	})
}

func infoHandler(cmd *cobra.Command, args []string) {
	var (
		info *ffmpeg.Media
		err  error
	)

	if info, err = ffmpeg.Probe(args[0]); err != nil {
		utils.Printf("<red>Error reading media:</reset> <red+hb>%s</reset>\n", err)
		return
	}

	lib.RenderMedia(info)
	lib.RenderStreams(info)
}
