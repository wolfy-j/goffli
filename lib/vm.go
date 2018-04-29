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

import "github.com/yuin/gopher-lua"

// VM wraps at top of lua VM and tmpDir storage.
type VM struct {
	l *lua.LState
	t tmpDir
}

// NewVM returns new instance of lua VM wrapper.
func NewVM(args []string) *VM {
	vm := &VM{
		l: lua.NewState(),
		t: newTmpDir(),
	}

	// additional functions and modules can be added here
	vm.l.SetGlobal("ask", vm.l.NewFunction(NewPrompter(args)))
	vm.l.SetGlobal("print", vm.l.NewFunction(Print))

	// additional modules
	vm.l.PreloadModule("ffmpeg", ffmpegModule)
	vm.l.PreloadModule("tmp", vm.t.luaModule)

	return vm
}

// DoFile executes given lua script by it's path.
func (vm *VM) DoFile(path string) error {
	return vm.l.DoFile(path)
}

// Close closes lua VM and cleans temp files (if any).
func (vm *VM) Close() {
	vm.l.Close()
	vm.t.Clean()
}
