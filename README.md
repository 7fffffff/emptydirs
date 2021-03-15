# emptydirs

Print a list of empty directories. Empty directories contain no files,
but may have empty directories.

```
Usage:
  emptydirs [OPTIONS] [ROOTS...]

Options:
  -0	separate empty directories with NUL, like find -print0
  -p	ignore permission errors
  -z	ignore zero length files

If no roots are provided, the current working directory is used
```

## Installation

After installing the [go compiler](https://golang.org/),

```
go install github.com/7fffffff/emptydirs@latest
```

## License

Blue Oak Model License 1.0.0  
<https://blueoakcouncil.org/license/1.0.0>
