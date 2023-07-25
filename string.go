package main

import "fmt"

const HorspoolMaxAlphabetSize = 500

func BruteForceStringMatch(text, pattern string) int {
	n, m := len(text), len(pattern)
	for i := 0; i <= n-m; i++ {
		j := 0
		for (j < m) && (pattern[j] == text[i+j]) {
			j++
		}
		if j == m {
			return i
		}
	}
	return -1
}

func ShiftTable(pattern string) []int {
	size, m := HorspoolMaxAlphabetSize, len(pattern)
	table := make([]int, size)
	// Populate the map with initial values
	for i := 0; i < size; i++ {
		table[i] = m
	}
	// Determine actual shift sizes
	for j := 0; j < m-1; j++ {
		table[pattern[j]] = m - 1 - j
	}
	return table
}

func HorspoolStringMatch(text, pattern string) int {
	n, m := len(text), len(pattern)
	table := ShiftTable(pattern)
	i := m - 1 // Position of the pattern's right end

	for i < n {
		k := 0 // Number of matched characters
		for (k <= m-1) && (pattern[m-1-k] == text[i-k]) {
			k++
		}
		if k == m {
			return i - m + 1
		} else {
			i += table[text[i]]
		}
	}
	return -1
}

func main() {
	text := "the_artic_sarcastic_barbaric_bar"
	pattern := "barbaric"
	bfMatch := BruteForceStringMatch(text, pattern)
	hpMatch := HorspoolStringMatch(text, pattern)
	fmt.Printf("Brute force match index: %d\n", bfMatch)
	fmt.Printf("Horspool match index: %d\n", hpMatch)
}
