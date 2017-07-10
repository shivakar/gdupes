[![Build Status](https://travis-ci.org/shivakar/gdupes.svg?branch=master)](https://travis-ci.org/shivakar/gdupes)

# gdupes
A multithreaded tool for identifying duplicate files written in Go.

gdupes is greatly inspired by [fdupes](https://github.com/adrianlopezroche/fdupes)

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
  -h, --help        Display this help message
  -m, --summarize   Summarize duplicates information
  -v, --version     Display gdupes versiongdu
```

## Contributing

Any and all contributions in the form of suggestions, feature requests, code reviews, bug reports and/or pull requests are greatly appreciated.

## License

gdupes is licensed under a MIT license.
