package gdupes

import (
	"io"
	"os"
	"strings"
	"sync"

	"github.com/shivakar/xxhash"
)

// hashFile returns hexadecimal string output of xxhash64 for the given file
func hashFile(filepath string) string {
	h := xxhash.NewXXHash64()
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	b := make([]byte, 1024*1024) // 1 MB buffer
	for {
		n, err := f.Read(b)

		if n > 0 {
			h.Write(b[:n])
		}
		if err == io.EOF {
			break
		}
	}
	return h.String()
}

// includeFileInOutput returns true if the file should be included as per
// the given configuration
func includeFileInOutput(c *Config, fm FileMeta,
	fms FileMetaSlice) (bool, int, error) {
	if !c.Hardlinks {
		// Not treating hardlinks as duplicates
		inc, idx, err := fms.ContainsInode(fm)
		return !inc, idx, err
	}
	return true, -1, nil
}

// ProcessFiles computes hashes and updates map of hashes for files to be
// processed
func ProcessFiles(c *Config, filesToProcess <-chan string,
	fileHashes map[string]FileMetaSlice,
	lock *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for f := range filesToProcess {
		h := hashFile(f)
		info, err := os.Stat(f)
		if err != nil {
			panic(err)
		}
		fm := FileMeta{Path: f, Info: info}
		lock.Lock()
		_, ok := fileHashes[h]
		if !ok {
			fileHashes[h] = []FileMeta{fm}
		} else {
			inc, idx, err := includeFileInOutput(c, fm, fileHashes[h])
			if err != nil {
				panic(err)
			}
			if !inc {
				// Check if this file has a longer filename. This is to ensure
				// that regardless of what order hardlinks are processed in, you
				// always get the same result
				if strings.Compare(fm.Path, fileHashes[h][idx].Path) == 1 {
					fileHashes[h][idx] = fm
				}
				lock.Unlock()
				continue
			}
			fileHashes[h] = append(fileHashes[h], fm)
		}
		lock.Unlock()
	}
}
