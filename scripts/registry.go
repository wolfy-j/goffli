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
	"fmt"
	"github.com/shibukawa/configdir"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Registry allows to add, remove and list all registered lua scripts.
type Registry struct {
	storage *configdir.Config
}

// NewRegistry returns new script directory based on app and vendor names.
func NewRegistry(vendor, application string) *Registry {
	return &Registry{
		storage: configdir.New(vendor, application).QueryCacheFolder(),
	}
}

// Get returns script by name or error
func (r *Registry) Get(name string) *Script {
	for _, script := range r.GetAll() {
		if script.Name == name {
			return script
		}
	}

	return nil
}

// Register new scripts script
func (r *Registry) Register(name, content string) error {
	file, err := r.storage.Create(fmt.Sprintf("%s.lua", name))
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(content)

	return nil
}

// Remove removes one goffli lua script.
func (r *Registry) Remove(name string) error {
	if !r.storage.Exists(fmt.Sprintf("%s.lua", name)) {
		return fmt.Errorf("no such script")
	}

	return os.Remove(filepath.Join(r.storage.Path, fmt.Sprintf("%s.lua", name)))
}

// GetAll returns list of registered scripts scripts associated with their aliases.
func (r *Registry) GetAll() []*Script {
	files, err := ioutil.ReadDir(r.storage.Path)
	if err != nil {
		return nil
	}

	scripts := make([]*Script, 0)
	for _, f := range files {
		script, err := NewScript(filepath.Join(r.storage.Path, f.Name()))
		if err != nil {
			fmt.Printf("error loading script %s: %s\n", f.Name(), err)
			continue
		}

		scripts = append(scripts, script)
	}

	return scripts
}
