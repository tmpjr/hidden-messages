package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var (
		text    string
		pattern string
		matches int
	)
	text = lines[0]
	pattern = lines[1]
	matches = patternCount(text, pattern)

	fmt.Printf("Pattern:(%s) was found %d times in the text input.\n", pattern, matches)
}

func patternCount(text string, pattern string) int {
	count := 0
	for i := 0; i < len(text)-len(pattern); i++ {
		if text[i:i+len(pattern)] == pattern {
			count++
		}
	}

	return count
}
