package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var (
		text  string
		kstr  string
		words string
	)
	text = lines[0]
	kstr = lines[1]
	k, err := strconv.ParseInt(kstr, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	words = frequentWords(text, k)

	fmt.Printf("Frequent words:[%s].\n", words)
}

func frequentWords(text string, k int64) string {
	var (
		i     int64
		max   int64
		words string
		limit int64
	)
	limit = int64(len(text)) - k
	patterns := make(map[string]int64)

	max = 0
	for i = 0; i < limit; i++ {
		w := text[i : i+k]
		patterns[w] += 1
		if patterns[w] > max {
			max = patterns[w]
		}
	}

	glue := ""
	for word, count := range patterns {
		if count == max {
			words += glue + word
			glue = " "
		}
	}

	return words
}
