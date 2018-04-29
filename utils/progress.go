package utils

import (
	"github.com/briandowns/spinner"
	"gopkg.in/cheggaaa/pb.v1"
	"time"
)

// Progress allows to represent processing state using progress bar or spinner.
type Progress struct {
	bar *pb.ProgressBar
	sp  *spinner.Spinner
}

// NewProgress returns progress bar or spinner based on given total duration.
func NewProgress(total time.Duration, forceSpinner bool) *Progress {
	p := &Progress{}

	if total == 0 || forceSpinner {
		p.sp = spinner.New(spinner.CharSets[43], 100*time.Millisecond)
		p.sp.Start()
	} else {
		p.bar = pb.New(int(total.Nanoseconds()))
		p.bar.Width = 100
		p.bar.ShowTimeLeft = true
		p.bar.SetUnits(pb.U_DURATION)
		p.bar.Start()
	}

	return p
}

// Set updates progress position based on current progress duration.
func (p *Progress) Set(t time.Duration) {
	if p.bar != nil {
		p.bar.Set(int(t.Nanoseconds()))
	}
}

// Finish finishes the progress bar or spinner.
func (p *Progress) Finish() {
	if p == nil {
		return
	}

	if p.bar != nil {
		p.bar.Finish()
	}

	if p.sp != nil {
		p.sp.Stop()
	}
}
