package finder

import (
	"io/fs"
	"path/filepath"
)

func Files(fsys fs.FS) int {
	var count int
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if filepath.Ext(p) == ".go" {
			count++
		}
		return nil
	})

	return count
}
