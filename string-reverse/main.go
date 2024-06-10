package main

import "fmt"

func main() {
	r := reverseString("hello world")
	fmt.Println(r)
}

// time complexity O(n)
// space complexity O(n)
func reverseString(s string) string {
	if len(s) < 1 {
		return s
	}

	chars := make([]string, len(s))
	for i, value := range s {
		chars[i] = string(value)
	}
	var result string
	for i := len(chars) - 1; i >= 0; i-- {
		result += chars[i]
	}

	return result
}
