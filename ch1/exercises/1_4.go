//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	allDupFiles := ""
	where := make(map[string]map[string]int)
	for _, file := range os.Args[1:] {
		count := make(map[string]int)
		b, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup4 error: %v\n", err)
			continue
		}
		scanner := bufio.NewScanner(bytes.NewReader(b))
		for scanner.Scan() {
			count[scanner.Text()]++
		}
		for _, dupCount := range count {
			if dupCount > 1 {
				where[file] = count
			}
		}
	}
	for file := range where {
		allDupFiles += file + " "
	}
	for _, countMap := range where {
		for line, count := range countMap {
			if count > 1 {
				fmt.Printf("%d %s\t(%s)\n", count, line, allDupFiles)
			}
		}
	}
}
