package commands

import (
	"context"
	"path/filepath"

	"github.com/andev0x/ctxd/internal/storage/sqlite"
	"github.com/andev0x/ctxd/internal/watcher"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch [path]",
	Short: "Watch a repository for changes and update the index incrementally",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		dbPath := filepath.Join(".ctxd", "graph.db")
		store, err := sqlite.NewStore(dbPath)
		if err != nil {
			return err
		}
		defer store.Close()

		w := watcher.NewWatcher(store, path)
		return w.Start(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
