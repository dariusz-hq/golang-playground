package main

import "fmt"

func main() {
	reverseString("hello world")
}

// time complexity O(n)
// space complexity O(n)
func reverseString(s string) {
	chars := make([]string, len(s))
	for i, value := range s {
		chars[i] = string(value)
	}
	var result string
	for i := len(chars) - 1; i >= 0; i-- {
		result += chars[i]
	}
	fmt.Println(result)
}
