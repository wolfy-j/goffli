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
