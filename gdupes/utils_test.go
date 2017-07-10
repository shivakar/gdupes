package gdupes_test

import (
	"testing"

	"github.com/shivakar/gdupes/gdupes"
	"github.com/stretchr/testify/assert"
)

func TestHumanizeSize(t *testing.T) {
	assert := assert.New(t)
	testdata := []struct {
		n        float64
		expected string
	}{
		{0, "0 B"},
		{999, "999 B"},
		{1000, "1 KB"},
		{2450, "2.45 KB"},
		{2459, "2.46 KB"},
		{4200000, "4.20 MB"},
		{9e9, "9 GB"},
		{10e12, "10 TB"},
		{1679e12, "1.68 PB"},
	}
	for _, t := range testdata {
		assert.Equal(t.expected, gdupes.HumanizeSize(t.n))
	}
}
