package gdupes

import (
	"io"
	"os"
	"sync"

	"github.com/shivakar/xxhash"
)

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

// ProcessFiles computes hashes and updates map of hashes for files to be
// processed
func ProcessFiles(filesToProcess <-chan string, fileHashes map[string][]string,
	lock *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for f := range filesToProcess {
		h := hashFile(f)
		lock.Lock()
		_, ok := fileHashes[h]
		if !ok {
			fileHashes[h] = []string{f}
		} else {
			fileHashes[h] = append(fileHashes[h], f)
		}
		lock.Unlock()
	}
}
