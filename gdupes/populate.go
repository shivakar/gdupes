package gdupes

import (
	"os"
	"path/filepath"
	"sync"
)

// includeFileForProcessing checks fileinfo against the requested configuration.
// Returns true if the file is to be included for processing, false otherwise
func includeFileForProcessing(c *Config, f os.FileInfo) bool {
	if c.NoEmpty && f.Size() == 0 {
		return false
	}
	// If not a regular file (this includes directories and symlinks)
	// don't add to the channel
	if !f.Mode().IsRegular() {
		return false
	}
	return true
}

// PopulateFiles recurses through the list of directories adding files to be
// processed
func PopulateFiles(c *Config, filesToProcess chan<- string, directories []string,
	wg *sync.WaitGroup) {
	defer wg.Done()
	for _, d := range directories {
		if c.Recurse {
			err := filepath.Walk(d, func(path string, info os.FileInfo,
				err error) error {
				if err != nil {
					return err
				}
				if includeFileForProcessing(c, info) {
					filesToProcess <- path
				}
				return nil
			})
			if err != nil {
				panic(err)
			}
		} else {
			fd, err := os.Open(d)
			if err != nil {
				panic(err)
			}
			files, err := fd.Readdir(-1)
			if err != nil {
				panic(err)
			}
			for _, f := range files {
				if includeFileForProcessing(c, f) {
					filesToProcess <- filepath.Join(d, f.Name())
				}
			}
		}
	}
	close(filesToProcess)
}
