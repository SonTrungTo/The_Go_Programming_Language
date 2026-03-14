//go:build ignore

/**
 * This program prints a duplicated line preceded by its count and the names of all files in which it appears.
 * Usage: Use with a list of named files.
 */

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	where := make(map[string]string)

	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "dup: no files\n")
		os.Exit(1)
	} else {
		for _, file := range files {
			b, err := os.ReadFile(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup err: %v\n", err)
				continue
			}
			input := bufio.NewScanner(bytes.NewReader(b))
			for input.Scan() {
				counts[input.Text()]++
				if where[input.Text()] == "" {
					where[input.Text()] = file
				} else if !strings.Contains(where[input.Text()], file) {
					where[input.Text()] += ", " + file
				}
			}
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\t\t%s\n", count, line, where[line])
		}
	}
}
