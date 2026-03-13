//go:build ignore

/**
 * This program counts the number of duplicated lines with args as named files,
 * using ReadFile
 * Usage: Use a list of named files.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			b, err := os.ReadFile(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup3 error: %v\n", err)
				continue
			}
			// // This is the better solution
			// scanner := bufio.NewScanner(bytes.NewReader(b))
			// for scanner.Scan() {
			// 	counts[scanner.Text()]++ // Remove the newline character from the end of the line
			// }
			lines := strings.Split(string(b), "\n")
			for _, line := range lines {
				counts[line]++
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
