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
	"os/exec"
	"time"
)

// Run runs ffmpeg with given set of arguments, optional callback will be used to report progress (current duration,
// total duration). Callback total duration can be 0 if unable to automatically detect.
func Run(args []string, callback func(c, t time.Duration)) error {
	cmd := exec.Command("ffmpeg", args...)

	if callback == nil {
		var cmdErr bytes.Buffer
		cmd.Stderr = &cmdErr

		if err := cmd.Run(); err != nil {
			return extractError(err, cmdErr.String())
		}

		return nil
	}

	// ffmpeg stdout is stderr
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	w, err := newWatcher(stderr, callback)
	if err != nil {
		return err
	}

	defer w.Close()

	if err := cmd.Run(); err != nil {
		return extractError(err, w.Output)
	}

	return nil
}
