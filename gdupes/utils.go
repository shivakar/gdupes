package gdupes

import (
	"fmt"
	"strings"
)

// HumanizeSize returns human readable representation of size in bytes
func HumanizeSize(n float64) string {
	out := ""
	switch {
	case n > 1e15:
		out = fmt.Sprintf("%.2f PB", n/1e15)
	case n > 1e12:
		out = fmt.Sprintf("%.2f TB", n/1e12)
	case n >= 1e9:
		out = fmt.Sprintf("%.2f GB", n/1e9)
	case n >= 1e6:
		out = fmt.Sprintf("%.2f MB", n/1e6)
	case n >= 1e3:
		out = fmt.Sprintf("%.2f KB", n/1e3)
	default:
		out = fmt.Sprintf("%.2f B", n)
	}
	out = strings.Replace(out, ".00", "", -1) // Cleaning up integer values

	return out
}
