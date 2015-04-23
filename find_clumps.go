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
	fmt.Printf("k: %d\nL:%d\nt:%d\nG:%d\n", k, L, t, len(genome))
}

func findClumps(genome string, k int64, L int64, t int64) string {
	var clumps string
	var i int64
	//patterns := make(map[string]int64)
	data := make(map[int64]map[string]int64)

	//glue := ""
	for i = 0; i < int64(len(genome))-L; i++ {
		m := make(map[string]int64)
		for j := i; j < L-k; j++ {
			//fmt.Printf("%d:%d\n", j, j+k)
			w := genome[j : j+k]
			m[w] += 1
			if m[w] >= t {
				data[i][w] = m[w]
			}
		}
		//fmt.Printf("l:%d::i:%d\n", len(genome), i)
	}

	fmt.Printf("%v\n", data)

	/*glue := ""
	for i = 0; i < L-k; i++ {
		w := genome[i : i+k]
		patterns[w] += 1
		if patterns[w] == t {
			clumps += glue + w
			glue = " "
		}
	}*/

	return clumps
}
