package main

import (
	"fmt"
)

// returns [0, 3, 4, 4, 6, 30, 31]
func main() {
	a1 := []int{0, 3, 4, 31}
	a2 := []int{4, 6, 30, 32, 33}
	r := mergeSortedArrays(a1, a2)
	fmt.Println(r)
}

func mergeSortedArrays(a1 []int, a2 []int) []int {
	if len(a1) == 0 {
		return a2
	}
	if len(a2) == 0 {
		return a1
	}

	var result []int

	// track items in both arrays while pushing the lower one to the resulting slice
	indexA1 := 0
	allA1Evaluated := false
	indexA2 := 0
	allA2Evaluated := false
	currentItemArray1 := a1[indexA1]
	currentItemArray2 := a2[indexA2]
	allEvaluated := false

	for !allEvaluated {
		fmt.Println("indexA1", indexA1, "indexA2", indexA2)
		fmt.Println("current1", currentItemArray1, "current2", currentItemArray2)

		if currentItemArray1 <= currentItemArray2 {
			result = append(result, currentItemArray1)
			// when item added, and it's the last one mark evaluation for a1 complete
			if indexA1 == len(a1)-1 {
				allA1Evaluated = true
			} else {
				indexA1++
				currentItemArray1 = a1[indexA1]
			}
		} else {
			result = append(result, currentItemArray2)
			// when item added, and it's the last one mark evaluation for a1 complete
			if indexA2 == len(a2)-1 {
				allA2Evaluated = true
			} else {
				indexA2++
				currentItemArray2 = a2[indexA2]
			}
		}

		// if we already evaluated elements of one array then just add rest items from the second one
		if allA1Evaluated {
			result = append(result, a2[indexA2:]...)
			allA2Evaluated = true
		}
		if allA2Evaluated {
			result = append(result, a1[indexA1:]...)
			allA1Evaluated = true
		}

		allEvaluated = allA1Evaluated && allA2Evaluated

		fmt.Println("result", result)
		fmt.Println(allA1Evaluated, allA2Evaluated)
	}

	return result
}
