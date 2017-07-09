package gdupes

import (
	"fmt"
	"sync"
	"time"
)

// Run orchestrates gdupes execution
func Run(c *Config) {
	st := time.Now()
	var wg sync.WaitGroup
	var lock sync.Mutex
	filesToProcess := make(chan string, 500)
	fileHashes := make(map[string][]string)

	// Populating filesToProcess
	wg.Add(1)
	go PopulateFiles(filesToProcess, c.Directories, &wg)

	// Creating workers to process
	wg.Add(c.NumWorkers)
	for i := 0; i < c.NumWorkers; i++ {
		go ProcessFiles(filesToProcess, fileHashes, &lock, &wg)
	}
	wg.Wait()

	nSets := 0
	nDups := 0

	for _, v := range fileHashes {
		if len(v) > 1 {
			nSets++
			nDups += len(v) - 1
		}
	}

	fmt.Printf("%d duplicate files (in %d sets)\n", nDups, nSets)
	fmt.Println("Total time for processing: ", time.Since(st))
}
