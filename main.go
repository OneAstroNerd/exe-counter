package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	path, err := filepath.Abs(os.Args[0])
	if err != nil {
		return
	}

	dir, filename := filepath.Split(path)
	nameParts := strings.Split(filename, ".")
	if len(nameParts) < 2 {
		return
	}

	numStr := nameParts[0]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}

	newFilename := strconv.Itoa(num+1) + "." + nameParts[1]
	newPath := filepath.Join(dir, newFilename)

	err = os.Rename(path, newPath)
	if err != nil {
	}
}
