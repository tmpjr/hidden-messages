package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tmpjr/stringutil"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var (
		pattern    string
		reverse    string
		complement string
	)
	pattern = lines[0]

	bases := map[string]string{"A": "T", "T": "A", "C": "G", "G": "C"}
	reverse = stringutil.Reverse(pattern)

	for i := 0; i < len(reverse); i++ {
		complement += bases[string(reverse[i])]
	}

	//fmt.Printf("\nP:%s\nR:%s\nC:%s\n", pattern, reverse, complement)
	fmt.Println(complement)
}
