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

	clumps := findClumps(genome, k, L, t)
	fmt.Println(clumps)
}

func findClumps(genome string, k int64, L int64, t int64) string {
	var (
		clumps string
		i      int64
		y      int64
	)

	data := make(map[string]int64)
	for i = 0; i < int64(len(genome))-L; i++ {
		text := genome[i : i+L]
		d := make(map[string]int64)
		for y = 0; y < int64(len(text))-k; y++ {
			w := text[y : y+k]
			d[w] += 1
			if d[w] >= t {
				data[w] = d[w]
			}
		}
	}

	glue := ""
	for clump := range data {
		clumps += glue + clump
		glue = " "
	}

	return clumps
}
