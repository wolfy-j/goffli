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
