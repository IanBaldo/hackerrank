// https://www.hackerrank.com/challenges/ctci-ransom-note
package main

import (
	"fmt"
)

func main () {
	// Reads Inputs
	size, _ := intScanln(2)
	magazine, _ := wordScan(size[0])
	ransom, _ := wordScan(size[1])

	// Creates a hashmap to count how many times a word occurs
	wordMap := make(map[string]int)
	for _, word := range magazine {
		wordMap[word]++
	}

	answer := "Yes"
	// Decreases the word count for each word in the ransom
	for _,word := range ransom {
		wordMap[word]--
		if wordMap[word] < 0 {
			answer = "No"
			break
		}
	}

	fmt.Println(answer)
}

// Scans "n" words into a string slice
func wordScan (n int) ([]string, error) {
	words := make([]string, n)
	y := make([]interface{}, n)

	for i:= range words {
		y[i] = &words[i]
	}

	n, err := fmt.Scanln(y...)
	words = words[:n]

	return words,err
}

// Scans "n" integers into an integer slice
func intScanln (n int ) ([]int, error) {
	x := make([]int, n)
	y := make([]interface{}, n)

	for i := range x {
		y[i] = &x[i]
	}

	n, err := fmt.Scanln(y...)
	x = x[:n]

	return x, err
}