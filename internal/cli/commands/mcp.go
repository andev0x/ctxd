package commands

import (
	"path/filepath"

	"github.com/PizenLabs/lea/internal/mcp"
	"github.com/PizenLabs/lea/internal/storage/sqlite"
	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Start the MCP server to expose lea to AI agents",
	Long:  `The mcp command starts a Model Context Protocol server over stdio, allowing AI agents like Claude or Pi to query the structural graph.`,
	RunE: func(_ *cobra.Command, _ []string) error {
		dbPath := filepath.Join(".lea", "graph.db")
		store, err := sqlite.NewStore(dbPath)
		if err != nil {
			return err
		}
		defer func() { _ = store.Close() }() 

		s := mcp.NewServer(store)
		return s.Start()
	},
}

func init() {
	rootCmd.AddCommand(mcpCmd)
}
