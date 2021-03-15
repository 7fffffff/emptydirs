// Licensed under the Blue Oak Model License 1.0.0:
// https://blueoakcouncil.org/license/1.0.0

package main

import (
	"path/filepath"
	"strings"
	"testing"
)

var expectedPaths = map[string]struct{}{
	"testdir/0":     {},
	"testdir/2/b/i": {},
	"testdir/2/c":   {},
}

func TestFindEmptyDirs(t *testing.T) {
	walker := &dirWalker{
		IgnoreEmptyFiles: true,
	}
	expectedFilepaths := map[string]struct{}{}
	for orig, _ := range expectedPaths {
		segments := strings.Split(orig, "/")
		expectedFilepaths[filepath.Join(segments...)] = struct{}{}
	}
	results := map[string]struct{}{}
	err := walker.FindEmptyDirs("testdir", func(emptyDirPath string) {
		results[emptyDirPath] = struct{}{}
	})
	if err != nil {
		t.Fatal(err)
	}
	for path, _ := range expectedFilepaths {
		if _, ok := results[path]; !ok {
			t.Errorf("missing: %s", path)
		}
	}
	for path, _ := range results {
		if _, ok := expectedFilepaths[path]; !ok {
			t.Errorf("unexpected: %s", path)
		}
	}
}
