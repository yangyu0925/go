package main

import "fmt"

func lengthOfNonrepeatingSubstr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch];
			ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfNonrepeatingSubstr("abcabcbb"))
}
