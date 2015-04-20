package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	text := string(bytes)
	pattern := "CTTGATCAT"
	occurrences := findOccurrences(text, pattern)
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
