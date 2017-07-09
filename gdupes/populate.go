package gdupes

import (
	"os"
	"path/filepath"
	"sync"
)

// PopulateFiles recurses through the list of directories adding files to be
// processed
func PopulateFiles(filesToProcess chan<- string, directories []string,
	wg *sync.WaitGroup) {
	defer wg.Done()
	for _, d := range directories {
		err := filepath.Walk(d, func(path string, info os.FileInfo,
			err error) error {
			if err != nil {
				return err
			}
			// If not a regular file (this includes directories and symlinks)
			// don't add to the channel
			if !info.Mode().IsRegular() {
				return nil
			}
			filesToProcess <- path
			return nil
		})
		if err != nil {
			panic(err)
		}
	}
	close(filesToProcess)
}
