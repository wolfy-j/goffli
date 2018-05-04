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
	"encoding/json"
	"github.com/yuin/gopher-lua"
	luajson "layeh.com/gopher-json"
)

// Encode marshals given interface into value within lua state (using json bridge)
func Encode(v interface{}, l *lua.LState) (lua.LValue, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	lv, err := luajson.Decode(l, data)
	if err != nil {
		return nil, err
	}

	return lv, nil
}

// Decode unmarshal lua value into given variable using json bridge (make sure to pass pointer!)
func Decode(lv lua.LValue, v interface{}) error {
	data, err := luajson.Encode(lv)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
