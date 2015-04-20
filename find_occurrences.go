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
		text        string
		pattern     string
		occurrences string
	)
	pattern = lines[0]
	text = lines[1]
	occurrences = findOccurrences(text, pattern)

	fmt.Printf("Occurrences:[%s]\n", occurrences)
}

func findOccurrences(text string, pattern string) string {
	var occurrences string
	glue := ""
	for i := 0; i < len(text)-len(pattern); i++ {
		if text[i:i+len(pattern)] == pattern {
			occurrences += fmt.Sprintf("%s%d", glue, i)
			glue = " "
		}
	}

	return occurrences
}
