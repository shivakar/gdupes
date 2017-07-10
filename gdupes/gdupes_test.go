package gdupes_test

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"testing"

	"github.com/shivakar/gdupes/gdupes"
	"github.com/stretchr/testify/assert"
)

func isStringSliceEqual(e []string, a []string) bool {
	if len(e) != len(a) {
		return false
	}
	for i, v := range e {
		if v != a[i] {
			return false
		}
	}
	return true
}

func isStringSSEqual(e [][]string, a [][]string) bool {
	for _, av := range a {
		found := false
		sort.Strings(av)
		for _, ev := range e {
			sort.Strings(ev)
			if isStringSliceEqual(ev, av) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func TestGdupes(t *testing.T) {
	c := &gdupes.Config{}
	dirs := []string{"testdata"}
	buf := &bytes.Buffer{}
	if err := os.Chdir(".."); err != nil {
		panic(err)
	}

	c.Writer = buf
	t.Run("default", func(t *testing.T) {
		expected := [][]string{
			{"testdata/zero.txt",
				"testdata/.hidden.txt",
			}, {
				"testdata/b_hardlink.txt",
				"testdata/b_copy.txt",
			}, {
				"testdata/a.txt",
				"testdata/a_copy.txt",
				"testdata/a_copy_copy.txt",
			},
		}

		assert := assert.New(t)
		dupfiles, err := gdupes.Run(c, dirs)
		assert.Nil(err)
		assert.True(isStringSSEqual(expected, dupfiles),
			"expected: %v,\ngot %v\n", expected, dupfiles)
	})

	c.PrintVersion = true
	buf.Reset()
	t.Run("--version", func(t *testing.T) {
		assert := assert.New(t)
		expected := fmt.Sprintf("gdupes v%s\n", gdupes.VERSION)
		gdupes.Run(c, dirs)
		assert.Equal(expected, buf.String())
	})
	c.PrintVersion = false

	c.Summarize = true
	buf.Reset()
	t.Run("--summarize", func(t *testing.T) {
		assert := assert.New(t)
		expected := "4 duplicate files (in 3 sets), occupying 12 B.\nTotal time for processing: "
		gdupes.Run(c, dirs)
		assert.Equal(expected, buf.String()[:len(expected)])
	})
	c.Summarize = false
}
