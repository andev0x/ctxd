// Package contracts defines parser interfaces used across the app.
package contracts

import (
	"context"
	graph "github.com/PizenLabs/lea/internal/graph/contracts"
)

// Parser defines the interface for language parsers that extract graph data.
type Parser interface {
	ParseFile(ctx context.Context, path string) ([]*graph.Node, []*graph.Edge, error)
	ExtractCalls(ctx context.Context, path string) ([]*graph.Edge, error)
	ExtractControlFlow(ctx context.Context, path string) ([]*graph.Edge, error)
}
