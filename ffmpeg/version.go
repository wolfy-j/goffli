package ffmpeg

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

// Version returns ffmpeg version in string form
func Version() (string, error) {
	cmd := exec.Command("ffmpeg", "-version")

	var cmdOut, cmdErr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &cmdOut, &cmdErr

	if err := cmd.Run(); err != nil {
		return "", extractError(err, cmdErr.String())
	}

	r, err := regexp.Compile(`ffmpeg version ([^ ]+)`)
	if err != nil {
		return "", err
	}

	match := r.FindStringSubmatch(cmdOut.String())
	if len(match) == 0 {
		return "", fmt.Errorf("undefined version")
	}

	return match[1], nil
}
