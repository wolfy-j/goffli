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

package scripts

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

// Script represents one pre-registered lua script.
type Script struct {
	// Name is short script name (filename without extension).
	Name string

	// Path is physical script locations
	Path string

	// Tags includes set of user defined tags associated with script body.
	Tags map[string]string
}

// String converts script definition into string.
func (s *Script) String() string {
	body, _ := json.MarshalIndent(s, "", "    ")
	return string(body)
}

// Tag gets tag value or returns placeholder.
func (s *Script) Tag(name, placeholder string) string {
	if value, ok := s.Tags[name]; ok {
		return value
	}

	return placeholder
}

// NewScript returns new Script instance based on given part of returns errors.
func NewScript(path string) (*Script, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	name := filepath.Base(path)
	name = name[0 : len(name)-len(filepath.Ext(name))]

	return &Script{
		Name: name,
		Path: path,
		Tags: fetchTags(body),
	}, nil
}

// fetch all tags defined in script body using `--@key: value` pattern. All tags are automatically lowercased.
func fetchTags(body []byte) (tags map[string]string) {
	r, err := regexp.Compile(`--@(.+):\s*(.*)`)
	if err != nil {
		panic(err)
	}

	tags = make(map[string]string)
	for _, match := range r.FindAllSubmatch(body, -1) {
		tags[strings.ToLower(string(match[1]))] = string(match[2])
	}

	return tags
}
