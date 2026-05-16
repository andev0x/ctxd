package contracts

type NodeType string

const (
	NodeFunction  NodeType = "function"
	NodeMethod    NodeType = "method"
	NodeStruct    NodeType = "struct"
	NodeInterface NodeType = "interface"
	NodePackage   NodeType = "package"
	NodeModule    NodeType = "module"
	NodeFlow      NodeType = "flow"
)

type Node struct {
	ID       string                 `json:"id"`
	Type     NodeType               `json:"type"`
	Name     string                 `json:"name"`
	File     string                 `json:"file"`
	Line     int                    `json:"line"`
	Metadata map[string]interface{} `json:"metadata"`
}
