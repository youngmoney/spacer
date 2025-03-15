package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrintBashError(s ...any) {
	PrintBash(append(append([]any{"echo"}, s...), ">&2; false")...)
}

func PrintBash(s ...any) {
	for i, p := range s {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(p)
	}
	fmt.Println()
}

func IsDir(path string) bool {
	p, e := filepath.EvalSymlinks(path)
	if e != nil {
		return false
	}
	fileInfo, err := os.Stat(p)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}
