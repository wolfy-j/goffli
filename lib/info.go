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
	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
	"github.com/wolfy-j/goffli/ffmpeg"
	"github.com/wolfy-j/goffli/utils"
	"os"
	"strings"
)

// RenderMedia displays basic information about media.
func RenderMedia(info *ffmpeg.Media) {
	utils.Printf("\n")

	tb := tablewriter.NewWriter(os.Stdout)
	tb.SetAutoWrapText(false)
	tb.SetBorder(false)
	tb.SetColumnSeparator(":")

	if tags, ok := info.Format["tags"].(map[string]interface{}); ok {
		if artist, ok := tags["artist"].(string); ok {
			tb.Append([]string{"Artist", utils.Sprintf("<red>%s</reset>", artist)})
		}

		if title, ok := tags["title"].(string); ok {
			tb.Append([]string{"Title", utils.Sprintf("<yellow>%s</reset>", title)})
		}

		if desc, ok := tags["description"].(string); ok {
			tb.Append([]string{"Description", utils.Sprintf("<yellow>%s</reset>", desc)})
		}
	}

	if format, ok := info.Format["format_long_name"].(string); ok {
		tb.Append([]string{"Format", utils.Sprintf("<green>%s</reset>", format)})
	}

	tb.Append([]string{"Size", utils.Sprintf("<cyan>%s</reset>", humanize.Bytes(info.Size))})

	if info.Duration != 0 {
		tb.Append([]string{"Duration", utils.Sprintf("<yellow>%s</reset>", utils.FormatDuration(info.Duration))})
	} else {
		tb.Append([]string{"Duration", utils.Sprintf("<red>undefined</reset>")})
	}

	tb.Render()
}

// RenderStreams displays table with streams information.
func RenderStreams(info *ffmpeg.Media) {
	utils.Printf("\n")

	tb := tablewriter.NewWriter(os.Stdout)
	tb.SetAutoWrapText(false)
	tb.SetHeader([]string{"Stream", "Label", "Format", "Details"})

	for _, stream := range info.Streams {
		tb.Append([]string{formatIndex(stream), formatLabel(stream), stream.CodecLongName, formatDetails(stream)})
	}

	tb.Render()
}

// formatIndex formats stream index including stream type and color coding.
func formatIndex(stream *ffmpeg.Stream) (out string) {
	var (
		cl = "white"
		tp = stream.Type
	)

	switch stream.Type {
	case "video":
		cl = "red"
	case "audio":
		cl = "yellow"
	case "subtitle":
		cl = "cyan"
		tp = "subtl"
	}

	return utils.Sprintf("%02v:<%s>%s</reset>", stream.Index, cl, tp)
}

// formatLabel fetches stream title
func formatLabel(info *ffmpeg.Stream) (label string) {
	if title, ok := info.Tags["title"]; ok {
		label += title
	}

	if title, ok := info.Tags["filename"]; ok {
		label += title
	}

	if label == "" {
		label = "-none-"
	}

	return
}

// formatDetails formats additional information about stream
func formatDetails(stream *ffmpeg.Stream) (out string) {
	if stream.Type == "video" {
		out = utils.Sprintf("<red>res</reset>: %vx%v", stream.Width, stream.Height)
	}

	if lang, ok := stream.Tags["language"]; ok {
		if lang != "und" {
			out += utils.Sprintf(", <green>lng</reset>: %s", lang)
		}
	}

	if out == "" {
		out = "-none-"
	}

	return strings.Trim(out, " ,")
}
