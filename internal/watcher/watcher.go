package watcher

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/andev0x/ctxd/internal/parser/golang"
	"github.com/andev0x/ctxd/internal/storage/contracts"
	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	store  contracts.Store
	parser *golang.Parser
	root   string
}

func NewWatcher(store contracts.Store, root string) *Watcher {
	return &Watcher{
		store:  store,
		parser: golang.NewParser(),
		root:   root,
	}
}

func (w *Watcher) Start(ctx context.Context) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	// Recursively add directories to watch
	err = filepath.Walk(w.root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if strings.HasPrefix(info.Name(), ".") && info.Name() != "." {
				return filepath.SkipDir
			}
			if info.Name() == "vendor" {
				return filepath.SkipDir
			}
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	fmt.Printf("Watching for changes in %s...\n", w.root)

	// Debounce timer map to avoid multiple rapid updates for the same file
	timers := make(map[string]*time.Timer)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}

			if !strings.HasSuffix(event.Name, ".go") {
				continue
			}

			if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
				if t, ok := timers[event.Name]; ok {
					t.Stop()
				}
				timers[event.Name] = time.AfterFunc(500*time.Millisecond, func() {
					w.handleUpdate(ctx, event.Name)
				})
			} else if event.Op&fsnotify.Remove == fsnotify.Remove || event.Op&fsnotify.Rename == fsnotify.Rename {
				w.handleDelete(ctx, event.Name)
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			log.Printf("error: %v", err)

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (w *Watcher) handleUpdate(ctx context.Context, path string) {
	fmt.Printf("File changed: %s, updating index...\n", path)
	
	// Surgical update: delete old nodes for this file and re-parse
	if err := w.store.DeleteByFile(ctx, path); err != nil {
		log.Printf("Error deleting old nodes for %s: %v", path, err)
		return
	}

	nodes, edges, err := w.parser.ParseFile(path)
	if err != nil {
		log.Printf("Error parsing %s: %v", path, err)
		return
	}

	for _, n := range nodes {
		if err := w.store.SaveNode(ctx, n); err != nil {
			log.Printf("Error saving node: %v", err)
		}
	}
	for _, e := range edges {
		if err := w.store.SaveEdge(ctx, e); err != nil {
			log.Printf("Error saving edge: %v", err)
		}
	}

	// Extract calls
	callEdges, _ := w.parser.ExtractCalls(path)
	for _, e := range callEdges {
		if err := w.store.SaveEdge(ctx, e); err != nil {
			log.Printf("Error saving call edge: %v", err)
		}
	}
	fmt.Printf("Updated %s\n", path)
}

func (w *Watcher) handleDelete(ctx context.Context, path string) {
	fmt.Printf("File deleted: %s, removing from index...\n", path)
	if err := w.store.DeleteByFile(ctx, path); err != nil {
		log.Printf("Error deleting nodes for %s: %v", path, err)
	}
}
