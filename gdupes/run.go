package gdupes

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"sync"
	"time"
)

// checkAddDirectory adds a directory to the list of directories if not already
// present
func checkAddDirectory(dirs []string, d string) []string {
	for _, v := range dirs {
		if v == d {
			return dirs
		}
	}
	return append(dirs, d)
}

// Run orchestrates gdupes execution
func Run(c *Config, args []string) ([][]string, error) {
	st := time.Now()
	if c.Writer == nil {
		c.Writer = os.Stdout
	}

	if c.PrintVersion {
		fmt.Fprintf(c.Writer, "gdupes v%s\n", VERSION)
		return nil, nil
	}
	if len(args) < 1 {
		return nil, errors.New("must specify at least one directory to scan")
	}
	for _, d := range args {
		if fi, err := os.Stat(d); err != nil || !fi.IsDir() {
			return nil, fmt.Errorf("directory '%s' does not exist", d)
		}
		c.Directories = checkAddDirectory(c.Directories, d)
	}

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

	suffix := "\n"
	if c.Sameline {
		suffix = " "
	}
	if !c.Summarize {
		for _, s := range out {
			for _, v := range s {
				fmt.Fprintf(c.Writer, "%s%s", v, suffix)
			}
			fmt.Fprintln(c.Writer)
		}
	}

	if c.Summarize {
		fmt.Fprintf(c.Writer, "%d duplicate files (in %d sets), occupying %s.\n",
			nDups, nSets, HumanizeSize(float64(tSize)))
		fmt.Fprintf(c.Writer, "Total time for processing: %v\n", time.Since(st))
	}

	return out, nil
}
