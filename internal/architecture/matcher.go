package architecture

import (
	"path/filepath"
	"strings"
)

// Matcher handles mapping files to architecture layers.
type Matcher struct {
	layers []Layer
}

// NewMatcher creates a new architecture layer matcher.
func NewMatcher(cfg *Config) *Matcher {
	return &Matcher{layers: cfg.Layers}
}

// LayerForFile determines which layer a file belongs to based on patterns.
func (m *Matcher) LayerForFile(path string) (string, bool) {
	if path == "" {
		return "", false
	}
	normalized := normalizePath(path)
	for _, layer := range m.layers {
		for _, pattern := range layer.Patterns {
			if matchPattern(normalized, normalizePath(pattern)) {
				return layer.Name, true
			}
		}
	}
	return "", false
}

func normalizePath(path string) string {
	cleaned := filepath.ToSlash(filepath.Clean(path))
	return strings.TrimPrefix(cleaned, "./")
}

func matchPattern(path, pattern string) bool {
	if pattern == "" {
		return false
	}
	if strings.HasSuffix(pattern, "/**") {
		base := strings.TrimSuffix(pattern, "/**")
		return strings.HasPrefix(path, base)
	}
	if strings.HasSuffix(pattern, "**") {
		base := strings.TrimSuffix(pattern, "**")
		return strings.HasPrefix(path, strings.TrimSuffix(base, "/"))
	}
	if ok, err := filepath.Match(pattern, path); err == nil && ok {
		return true
	}
	return strings.Contains(path, pattern)
}
