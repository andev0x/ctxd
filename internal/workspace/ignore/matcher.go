package ignore

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// Matcher defines simple ignore rules for directories and files during indexing.
type Matcher struct {
	root       string
	dirNames   map[string]struct{}
	fileNames  map[string]struct{}
	extensions map[string]struct{}
}

func NewMatcher(root string) *Matcher {
	return &Matcher{
		root: root,
		dirNames: map[string]struct{}{
			".lea":        {},
			".git":         {},
			".idea":        {},
			".vscode":      {},
			"bin":          {},
			"build":        {},
			"dist":         {},
			"node_modules": {},
			"target":       {},
			"vendor":       {},
		},
		fileNames: map[string]struct{}{
			".DS_Store": {},
		},
		extensions: map[string]struct{}{
			".exe": {},
			".bin": {},
			".so":  {},
		},
	}
}

func (m *Matcher) ShouldSkipDir(path string, entry fs.DirEntry) bool {
	if path == m.root {
		return false
	}
	name := entry.Name()
	if strings.HasPrefix(name, ".") {
		return true
	}
	_, ok := m.dirNames[name]
	return ok
}

func (m *Matcher) ShouldSkipFile(path string, entry fs.DirEntry) bool {
	name := entry.Name()
	if strings.HasPrefix(name, ".") {
		return true
	}
	if _, ok := m.fileNames[name]; ok {
		return true
	}
	if _, ok := m.extensions[strings.ToLower(filepath.Ext(name))]; ok {
		return true
	}
	return false
}
