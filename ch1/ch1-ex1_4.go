/* Dup2 prints the count and text of lines that appear
more than once in the input.
It reads from stdin or from a list of named files.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:] // array of strings
	for _, filename := range files {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts, filename)
		f.Close()
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

/* countLines: function which takes a file descriptor,
and a map/dictionary and modifies the dictionary by
adding each line as a keyword and the number of duplicates
of each line as the associated value */
func countLines(f *os.File, counts map[string]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		s := fmt.Sprintf("%s: %s", filename, input.Text())
		counts[s]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
