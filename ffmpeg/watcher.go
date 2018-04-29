package ffmpeg

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type watcher struct {
	Output   string
	scanner  *bufio.Scanner
	updates  chan string
	done     chan interface{}
	callback func(current, total time.Duration)
	rd, rp   *regexp.Regexp
}

// newWatcher watches for the Output and updates callback (attention, total duration can be 0 if ffmpeg unable to detect it).
func newWatcher(pipe io.ReadCloser, progress func(c, t time.Duration)) (w *watcher, err error) {
	w = &watcher{
		callback: progress,
		updates:  make(chan string),
		done:     make(chan interface{}),
	}

	if w.rd, err = regexp.Compile(`Duration:\s*([0-9.\:]+)`); err != nil {
		return nil, err
	}

	if w.rp, err = regexp.Compile(`time=\s*([0-9.\:]+)`); err != nil {
		return nil, err
	}

	w.scanner = bufio.NewScanner(pipe)
	w.scanner.Split(w.handle)

	go w.scanner.Scan()
	go w.watch()

	return w, nil
}

func (w *watcher) watch() {
	var total time.Duration

	for {
		select {
		case update := <-w.updates:
			if update != "" {
				if total == 0 {
					fd := w.rd.FindStringSubmatch(update)
					if len(fd) != 0 {
						total = parseDuration(fd[1])
					}
				}

				ft := w.rp.FindAllStringSubmatch(update, -1)
				if len(ft) > 1 {
					w.callback(parseDuration(ft[len(ft)-1][1]), total)
				}
			}

			w.Output += update
			update = ""
		case <-w.done:
			return
		}
	}
}

func (w *watcher) handle(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// pushing updates update
	w.updates <- string(data)

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}

	// Request more updates.
	return 0, nil, nil
}

func (w *watcher) Close() {
	close(w.done)
}

// format 00:00:14.20
func parseDuration(framePosition string) (duration time.Duration) {
	segments := strings.Split(framePosition, ":")

	if hours, err := strconv.Atoi(segments[0]); err == nil {
		duration += time.Hour * time.Duration(hours)
	}

	if minutes, err := strconv.Atoi(segments[1]); err == nil {
		duration += time.Minute * time.Duration(minutes)
	}

	if seconds, err := strconv.ParseFloat(segments[2], 64); err == nil {
		duration += time.Second * time.Duration(seconds)
	}

	return
}
