package gdupes_test

import (
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
	if err := os.Chdir(".."); err != nil {
		panic(err)
	}

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

	// c.Recurse = true
	// t.Run("--recurse", func(t *testing.T) {
	// 	fmt.Println("Inside --recurse", c.Recurse)
	// })
	// c.Recurse = false
	// t.Run("--hardlinks", func(t *testing.T) {
	// 	fmt.Println("Inside --hardlinks", c.Recurse)
	// })
}
