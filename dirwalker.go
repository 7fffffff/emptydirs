// Licensed under the Blue Oak Model License 1.0.0:
// https://blueoakcouncil.org/license/1.0.0

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type dirWalker struct {
	IgnoreEmptyFiles       bool
	IgnorePermissionErrors bool
}

func (walker *dirWalker) FindEmptyDirs(root string, onEmptyDir func(emptyDirPath string)) error {
	root = filepath.Clean(root)
	allDirs := []string{}
	emptyDirs := map[string]struct{}{}
	nonEmptyDirs := map[string]struct{}{}
	// If there's a file, then by definition all of the parent dirs are non-empty
	nonEmptyPathFound := func(path string) {
		dir := filepath.Dir(path)
		if _, ok := nonEmptyDirs[dir]; ok {
			return
		}
		for dir != root && dir != "." && dir != string(filepath.Separator) {
			delete(emptyDirs, dir)
			nonEmptyDirs[dir] = struct{}{}
			dir = filepath.Dir(dir)
		}
	}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// When ignoring permission errors, treat the path as non-empty
		if os.IsPermission(err) && walker.IgnorePermissionErrors {
			fmt.Fprintln(os.Stderr, err)
			nonEmptyPathFound(path)
			return nil
		}
		if err != nil {
			return err
		}
		if info.IsDir() {
			if path != root {
				allDirs = append(allDirs, path)
				emptyDirs[path] = struct{}{}
			}
			return nil
		}
		if walker.IgnoreEmptyFiles && info.Size() == 0 {
			return nil
		}
		nonEmptyPathFound(path)
		return nil
	})
	if err != nil {
		return err
	}
	for _, dir := range allDirs {
		parent := filepath.Dir(dir)
		if _, ok := emptyDirs[parent]; ok {
			continue
		}
		if _, ok := emptyDirs[dir]; ok {
			onEmptyDir(dir)
		}
	}
	return nil
}
