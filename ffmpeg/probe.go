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
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"strconv"
)

// Media describes file information.
type Media struct {
	// Format common file format information.
	Format map[string]interface{} `json:"format"`

	// Streams item information structure.
	Streams []*Stream `json:"streams"`

	// Duration is duration of storage file in seconds.
	Duration float64

	// Size of file in bytes.
	Size uint64
}

// Strings packs info into json.
func (info *Media) String() string {
	body, _ := json.MarshalIndent(info, "", "  ")
	return string(body)
}

// Stream information about single stream.
type Stream struct {
	Index         int               `json:"index"`
	Type          string            `json:"codec_type"`
	CodecName     string            `json:"codec_name"`
	CodecLongName string            `json:"codec_long_name"`
	Bitrate       string            `json:"bitRate"`
	Duration      float64           `json:"duration,string"`
	Width         int64             `json:"width"`
	Height        int64             `json:"height"`
	Tags          map[string]string `json:"tags,omitempty"`
}

// Strings packs info into json.
func (info *Stream) String() string {
	body, _ := json.MarshalIndent(info, "", "  ")
	return string(body)
}

// Probe returns container and stream information from given file.
func Probe(filename string) (*Media, error) {
	stat, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(
		"ffprobe",
		"-v",
		"quiet",
		"-print_format",
		"json",
		"-show_format",
		"-show_streams",
		filename,
	)

	var cmdOut, cmdErr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &cmdOut, &cmdErr

	if err := cmd.Run(); err != nil {
		return nil, extractError(err, cmdErr.String())
	}

	info := &Media{}
	if err := json.Unmarshal([]byte(cmdOut.String()), &info); err != nil {
		return nil, err
	}

	// post processing
	streams := make([]*Stream, 0)
	for _, stream := range info.Streams {
		if stream.Type == "video" {
			if stream.Width == 0 || stream.Height == 0 {
				continue
			}
		}

		if info.Duration < stream.Duration {
			info.Duration = stream.Duration
		}

		streams = append(streams, stream)
	}

	info.Streams = streams

	// fallback duration using container information
	if info.Duration == 0 {
		if _, ok := info.Format["duration"]; ok {
			if d, err := strconv.ParseFloat(info.Format["duration"].(string), 64); err != nil {
				info.Duration = float64(d)
			}
		}
	}

	if size, ok := info.Format["size"]; ok {
		if intSize, err := strconv.Atoi(size.(string)); err != nil {
			info.Size = uint64(intSize)
		} else {
			info.Size = uint64(stat.Size())
		}
	}

	for _, stream := range info.Streams {
		if stream.Duration == 0 {
			stream.Duration = info.Duration
		}

		if stream.Bitrate == "" {
			stream.Bitrate = "0"
		}

		// flipping width and height
		if angle, rotate := stream.Tags["rotate"]; rotate {
			if angle == "90" || angle == "-90" {
				// please do not rotate by more than 180
				stream.Width, stream.Height = stream.Height, stream.Width
			}
		}
	}

	return info, nil
}
