// Licensed under the Blue Oak Model License 1.0.0:
// https://blueoakcouncil.org/license/1.0.0

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	print0 := false
	walker := &dirWalker{}
	flagset := flag.NewFlagSet("emptydirs", flag.ExitOnError)
	flagset.BoolVar(&print0, "0", false, "separate empty directories with NUL, like find -print0")
	flagset.BoolVar(&walker.IgnorePermissionErrors, "p", false, "ignore permission errors")
	flagset.BoolVar(&walker.IgnoreEmptyFiles, "z", false, "ignore empty files")
	flagset.Usage = func() {
		fmt.Println("Print a list of empty directories. Empty directories contain no files,")
		fmt.Println("but may have empty directories.")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  emptydirs [OPTIONS] [ROOTS...]")
		fmt.Println()
		fmt.Println("Options:")
		flagset.PrintDefaults()
		fmt.Println()
		fmt.Println("If no roots are provided, the current working directory is used")
	}
	err := flagset.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	roots := flagset.Args()
	if len(roots) == 0 {
		roots = append(roots, ".")
	}
	for _, root := range roots {
		err := walker.FindEmptyDirs(root, func(emptyDirPath string) {
			if print0 {
				os.Stdout.WriteString(emptyDirPath)
				os.Stdout.Write([]byte{0})
				return
			}
			fmt.Println(emptyDirPath)
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
