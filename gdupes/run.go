package gdupes

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"sort"
	"sync"
	"time"
)

// Run orchestrates gdupes execution
func Run(c *Config, args []string) ([][]string, error) {
	st := time.Now()

	if c.PrintVersion {
		fmt.Printf("gdupes v%s\n", VERSION)
		os.Exit(0)
	}
	if len(args) < 1 {
		return nil, errors.New("must specify at least one directory to scan")
	}
	for _, d := range args {
		if fi, err := os.Stat(d); err != nil || !fi.IsDir() {
			return nil, fmt.Errorf("directory '%s' does not exist", d)
		}
		c.Directories = append(c.Directories, d)
	}
	c.NumWorkers = 2 * runtime.NumCPU()

	var wg sync.WaitGroup
	var lock sync.Mutex

	filesToProcess := make(chan string, 500)
	fileHashes := make(map[string]FileMetaSlice)

	// Populating filesToProcess
	wg.Add(1)
	go PopulateFiles(c, filesToProcess, c.Directories, &wg)

	// Creating workers to process
	wg.Add(c.NumWorkers)
	for i := 0; i < c.NumWorkers; i++ {
		go ProcessFiles(c, filesToProcess, fileHashes, &lock, &wg)
	}
	wg.Wait()

	for _, s := range fileHashes {
		if len(s) > 1 {
			for _, v := range s {
				fmt.Println(v.Path)
			}
			fmt.Println()
		}
	}

	nSets := 0
	nDups := 0
	tSize := int64(0)

	out := make([][]string, 0)
	for _, v := range fileHashes {
		if len(v) > 1 {
			nSets++
			nDups += len(v) - 1
			for i, fm := range v {
				if i > 0 {
					tSize += fm.Info.Size()
				}
			}
			duplicates := v.GetFilenames()
			sort.Strings(duplicates)
			out = append(out, duplicates)
		}
	}

	fmt.Printf("%d duplicate files (in %d sets), occupying %s\n",
		nDups, nSets, HumanizeSize(float64(tSize)))
	fmt.Println("Total time for processing: ", time.Since(st))

	return out, nil
}
