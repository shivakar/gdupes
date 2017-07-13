// Copyright © 2017 Shivakar Vulli <svulli@shivakar.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/shivakar/gdupes/gdupes"
	"github.com/spf13/cobra"
)

var config gdupes.Config

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gdupes DIRECTORY...",
	Short: "A multithreaded CLI tool for identifying duplicate files",
	Long:  `gdupes is a multithreaded command-line tool for identifying duplicate files.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := gdupes.Run(&config, args)
		return err
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.Flags().BoolVarP(&config.Recurse, "recurse", "r", false, "Recurse through subdirectories")
	// RootCmd.Flags().BoolVarP(&config.Symlinks, "symlinks", "s", false, "Follow symlinks")
	RootCmd.Flags().BoolVarP(&config.Hardlinks, "hardlinks", "H", false, "Treat hardlinks as duplicates")
	RootCmd.Flags().BoolVarP(&config.NoEmpty, "noempty", "n", false, "Exclude zero-length/empty files")
	RootCmd.Flags().BoolVarP(&config.NoHidden, "nohidden", "A", false, "Exclude hidden files (POSIX only)")
	RootCmd.Flags().BoolVarP(&config.Sameline, "sameline", "1", false, "List set of matches on the same line")
	// RootCmd.Flags().BoolVarP(&config.Size, "size", "S", true, "Show size of duplicate files")
	RootCmd.Flags().BoolVarP(&config.Summarize, "summarize", "m", false, "Summarize duplicates information")
	// RootCmd.Flags().BoolVarP(&config.Quiet, "quiet", "q", false, "Hide progress indicator")
	RootCmd.Flags().BoolVarP(&config.PrintVersion, "version", "v", false, "Display gdupes version")
	RootCmd.Flags().BoolP("help", "h", false, "Display this help message")
}
