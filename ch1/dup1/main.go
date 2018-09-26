// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create map of string/int combinations
	counts := make(map[string]int)
	// Read input & break it into lines/workds
	input := bufio.NewScanner(os.Stdin)
	// Reads next line and removes newline character, result is input.Text()
	// returns true if there is a line and false where no more input
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			// Returns formatted output from a list
			// First argument is a format string  using conversion characters (letter%)
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-
