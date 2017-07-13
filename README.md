[![Build Status](https://travis-ci.org/shivakar/gdupes.svg?branch=master)](https://travis-ci.org/shivakar/gdupes)

# gdupes
A multithreaded tool for identifying duplicate files written in Go.

gdupes CLI is greatly inspired by [fdupes](https://github.com/adrianlopezroche/fdupes). One of the design goals was to keep the option flags and the tool output as close to fdupes as possible to be a viable drop-in replacement.

## Installation

```
go get -u github.com/shivakar/gdupes
```

## Usage

```
gdupes [flags] DIRECTORY...
```

To see full help message, `gdupes --help`:

```
gdupes is a multithreaded command-line tool for identifying duplicate files.

Usage:
  gdupes DIRECTORY... [flags]

Flags:
  -H, --hardlinks   Treat hardlinks as duplicates
  -h, --help        Display this help message
  -n, --noempty     Exclude zero-length/empty files
  -A, --nohidden    Exclude hidden files (POSIX only)
  -r, --recurse     Recurse through subdirectories
  -1, --sameline    List set of matches on the same line
  -m, --summarize   Summarize duplicates information
  -v, --version     Display gdupes version
```

## Contributing

Any and all contributions in the form of suggestions, feature requests, code reviews, bug reports and/or pull requests are greatly appreciated.

## License

gdupes is licensed under a MIT license.
