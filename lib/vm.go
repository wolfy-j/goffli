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
