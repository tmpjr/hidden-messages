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
		pattern    string
		reverse    string
		complement string
	)
	pattern = lines[0]

	bases := map[string]string{"A": "T", "T": "A", "C": "G", "G": "C"}

	runes := []rune(pattern)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reverse = string(runes)

	for i := 0; i < len(reverse); i++ {
		complement += bases[string(reverse[i])]
	}

	//fmt.Printf("\nP:%s\nR:%s\nC:%s\n", pattern, reverse, complement)
	fmt.Println(complement)
}
