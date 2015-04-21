/*
 * Clump Finding Problem: Find patterns forming clumps in a string.
 *    Input: A string Genome, and integers k, L, and t.
 *    Output: All distinct k-mers forming (L, t)-clumps in Genome.
 *
 * Sample Input:
 *	 CGGACTCGACAGATGTGAAGAACGACAATGTGAAGACTCGACACGACAGAGTGAAGAGAAGAGGAAACATTGTAA
 *   5 50 4
 *
 * Sample Output:
 *	 CGACA GAAGA
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	contents := string(bytes)

	lines := strings.Split(contents, "\n")
	parts := strings.Split(lines[1], " ")
	genome := lines[0]
	k, _ := strconv.ParseInt(parts[0], 0, 0)
	L, _ := strconv.ParseInt(parts[1], 0, 0)
	t, _ := strconv.ParseInt(parts[2], 0, 0)

	fmt.Printf("%s\n%d %d %d\n", genome, k, L, t)
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
