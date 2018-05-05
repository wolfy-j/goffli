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
func NewProgress(total time.Duration, pbType string) *Progress {
	p := &Progress{}

	if total == 0 || pbType == "spinner" {
		p.sp = spinner.New(spinner.CharSets[43], 100*time.Millisecond)
		p.sp.Start()
	} else if pbType != "none" {
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
	if p.bar != nil {
		p.bar.Finish()
	}

	if p.sp != nil {
		p.sp.Stop()
	}
}
