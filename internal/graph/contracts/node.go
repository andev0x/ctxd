package contracts

// NodeType represents the kind of structural element a node represents.
type NodeType string

const (
	// NodeFunction represents a standalone function.
	NodeFunction NodeType = "function"
	// NodeMethod represents a method belonging to a type.
	NodeMethod NodeType = "method"
	// NodeStruct represents a structure or class definition.
	NodeStruct NodeType = "struct"
	// NodeInterface represents an interface or protocol definition.
	NodeInterface NodeType = "interface"
	// NodePackage represents a software package or namespace.
	NodePackage NodeType = "package"
	// NodeModule represents a software module.
	NodeModule NodeType = "module"
	// NodeFlow represents a data or control flow path.
	NodeFlow NodeType = "flow"
)

// Node represents a structural element (symbol) in the codebase.
type Node struct {
	ID       string                 `json:"id"`
	Type     NodeType               `json:"type"`
	Name     string                 `json:"name"`
	File     string                 `json:"file"`
	Line     int                    `json:"line"`
	Metadata map[string]interface{} `json:"metadata"`
}
