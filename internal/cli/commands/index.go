package commands

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/andev0x/ctxd/internal/parser/golang"
	"github.com/andev0x/ctxd/internal/storage/sqlite"
	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:   "index [path]",
	Short: "Index a repository",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		// Ensure .ctxd directory exists
		ctxdDir := filepath.Join(path, ".ctxd")
		if _, err := os.Stat(ctxdDir); os.IsNotExist(err) {
			if err := os.Mkdir(ctxdDir, 0755); err != nil {
				return fmt.Errorf("failed to create .ctxd directory: %w", err)
			}
		}

		dbPath := filepath.Join(ctxdDir, "graph.db")
		store, err := sqlite.NewStore(dbPath)
		if err != nil {
			return err
		}
		defer store.Close()

		p := golang.NewParser()
		ctx := context.Background()

		err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				// Skip hidden directories (except current dir)
				if strings.HasPrefix(info.Name(), ".") && info.Name() != "." && info.Name() != ".ctxd" {
					return filepath.SkipDir
				}
				// Skip vendor
				if info.Name() == "vendor" {
					return filepath.SkipDir
				}
				return nil
			}

			if !strings.HasSuffix(filePath, ".go") {
				return nil
			}

			fmt.Printf("Indexing %s...\n", filePath)
			nodes, edges, err := p.ParseFile(filePath)
			if err != nil {
				fmt.Printf("Warning: failed to parse %s: %v\n", filePath, err)
				return nil
			}

			for _, n := range nodes {
				if err := store.SaveNode(ctx, n); err != nil {
					return err
				}
			}
			for _, e := range edges {
				if err := store.SaveEdge(ctx, e); err != nil {
					return err
				}
			}

			// Also extract calls
			callEdges, _ := p.ExtractCalls(filePath)
			for _, e := range callEdges {
				if err := store.SaveEdge(ctx, e); err != nil {
					return err
				}
			}

			return nil
		})

		if err != nil {
			return err
		}

		fmt.Println("Indexing complete.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)
}
