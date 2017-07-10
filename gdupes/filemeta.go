package gdupes

import (
	"errors"
	"os"
	"syscall"
)

// getStatStruct returns the syscall.Stat_t struct underlying FileInfo
func getStatStruct(fi os.FileInfo) (*syscall.Stat_t, error) {
	s, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return nil, errors.New("conversion to *syscall.Stat_t failed")
	}
	return s, nil
}

// FileMeta struct represents metadata for a file
type FileMeta struct {
	Path string
	Info os.FileInfo
}

// FileMetaSlice represents a FileMeta slice
type FileMetaSlice []FileMeta

// ContainsInode checks if the inode pointed to by the FileMeta is already
// contained in the FiletMetaSlice
func (f FileMetaSlice) ContainsInode(fm FileMeta) (bool, error) {
	s, err := getStatStruct(fm.Info)
	if err != nil {
		return false, err
	}
	if s.Nlink == 1 {
		return false, nil
	}
	inode := s.Ino
	for _, fim := range f {
		s, err := getStatStruct(fim.Info)
		if err != nil {
			return false, err
		}
		if s.Ino == inode {
			return true, nil
		}
	}
	return false, nil
}

// GetFilenames returns the list of filenames
func (f FileMetaSlice) GetFilenames() []string {
	out := make([]string, 0, len(f))
	for _, v := range f {
		out = append(out, v.Path)
	}
	return out
}
