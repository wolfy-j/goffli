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
	"fmt"
	"time"
)

// FormatDuration renders human friendly duration string.
func FormatDuration(d interface{}) string {
	var t time.Duration
	switch d.(type) {
	case float64:
		t = time.Second * time.Duration(d.(float64))
	case int64:
		t = time.Second * time.Duration(d.(int64))
	case int:
		t = time.Second * time.Duration(d.(int))
	case time.Duration:
		t = d.(time.Duration)
	default:
		return fmt.Sprintf("%v", d)
	}

	h := int(t.Hours())
	i := int(t.Minutes()) % 60
	s := t.Seconds() - float64(h*3600+i*60)

	return fmt.Sprintf("%02d:%02d:%02.0f", h, i, s)
}
