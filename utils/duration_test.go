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
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Duration_Time(t *testing.T) {
	assert.Equal(t, FormatDuration(time.Second), "00:00:01")
	assert.Equal(t, FormatDuration(time.Minute), "00:01:00")
	assert.Equal(t, FormatDuration(time.Hour), "01:00:00")
	assert.Equal(t, FormatDuration(45*time.Hour+61*time.Minute+30*time.Second), "46:01:30")
}

func Test_Duration_Float(t *testing.T) {
	assert.Equal(t, FormatDuration(1.0), "00:00:01")
	assert.Equal(t, FormatDuration(60.0), "00:01:00")
	assert.Equal(t, FormatDuration(3600.0), "01:00:00")
	assert.Equal(t, FormatDuration(3600.0*45+61*60.0+30.0), "46:01:30")
}

func Test_Duration_Int(t *testing.T) {
	assert.Equal(t, FormatDuration(1), "00:00:01")
	assert.Equal(t, FormatDuration(60), "00:01:00")
	assert.Equal(t, FormatDuration(3600), "01:00:00")
	assert.Equal(t, FormatDuration(3600*45+61*60+30), "46:01:30")
}

func Test_Duration_Fallback(t *testing.T) {
	assert.Equal(t, FormatDuration("value"), "value")
}
