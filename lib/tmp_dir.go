package lib

import (
	"fmt"
	"git.spiralscout.com/wolfy-j/goffli/utils"
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

// file returns new temp file name withing temp directory.
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
