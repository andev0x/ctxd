// Package python provides tree-sitter queries for Python.
package python

// SymbolsQuery is the tree-sitter query used to extract symbols from Python files.
const (
	SymbolsQuery = `
(function_definition
  name: (identifier) @function.name) @function.def

(class_definition
  name: (identifier) @class.name) @class.def
`
)
