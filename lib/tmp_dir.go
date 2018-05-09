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
	"fmt"
	"github.com/wolfy-j/goffli/utils"
	"github.com/yuin/gopher-lua"
	"io/ioutil"
	"os"
)

type tmpDir string

// newTmpDir allocates new temporary directory
func newTmpDir() tmpDir {
	dir, err := ioutil.TempDir(os.TempDir(), "goffli")
	if err != nil {
		panic(err)
	}

	return tmpDir(dir)
}

// luaModule returns new t module initializer
func (t tmpDir) luaModule(l *lua.LState) int {
	l.Push(l.SetFuncs(l.NewTable(), map[string]lua.LGFunction{
		"dir":  t.dir,
		"file": t.file,
	}))

	return 1
}

// dir returns temporary dir associated with the process.
func (t tmpDir) dir(l *lua.LState) int {
	l.Push(lua.LString(string(t)))
	return 1
}

// file returns new temp file name within temp directory.
func (t tmpDir) file(l *lua.LState) int {
	ext := l.OptString(1, "tmp")

	path := ""
	for i := 0; i < 100; i++ {
		path = fmt.Sprintf("%s%s%s.%s", string(t), string(os.PathSeparator), utils.RandString(10), ext)

		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
		if os.IsExist(err) {
			path = ""
			continue
		} else {
			f.Close()
			break
		}
	}

	if path == "" {
		return 0
	}

	utils.Log("tmp", path)
	l.Push(lua.LString(path))

	return 1
}

// Clean removes all temp files from tmpDir dir.
func (t tmpDir) Clean() {
	os.RemoveAll(string(t))
}
