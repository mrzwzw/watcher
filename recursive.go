package watcher

import (
	"os"
)

type Recursive interface {
	ListRecursive(name string, ignored map[string]struct{},
		ignoreHidden bool, fns []FilterFileHookFunc) (map[string]os.FileInfo, error)
}
