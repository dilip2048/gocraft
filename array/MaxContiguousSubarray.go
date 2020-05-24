package main

import (
	"fmt"
	"math"
)

// method to print largest contiguous array sum
func findMaxSubarray(a []int) (int, int) {
	start := 0
	s := 0
	end := 0
	currentMax := 0
	maxSoFar := math.MinInt64
	for i := 0; i < len(a); i++ {
		currentMax = a[i] + currentMax
		if currentMax > maxSoFar {
			maxSoFar = currentMax
			start = s
			end = i
		}
		if currentMax < 0 {
			currentMax = 0
			s = i + 1
		}
	}
	return start, end
}

func main() {
	array := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	start, end := findMaxSubarray(array)
	fmt.Println(array[start : end+1])
}
