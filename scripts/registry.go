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
