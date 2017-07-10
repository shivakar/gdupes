package gdupes

import "io"

const VERSION = "0.1.0"

// Config structure to store
type Config struct {
	Directories  []string
	Recurse      bool
	Symlinks     bool
	Hardlinks    bool
	NoEmpty      bool
	NoHidden     bool
	Sameline     bool
	Size         bool
	Summarize    bool
	Quiet        bool
	PrintVersion bool
	NumWorkers   int
	Writer       io.Writer
}
