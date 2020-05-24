package main

import "k8s.io/utils/integer"

// method to print largest contiguous array sum
func findMaxSum(a []int) int {
	currMax := a[0]
	maxSofar := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] + currMax > a[i]{
			currMax = a[i] + currMax
		}else {
			currMax = a[i]
		}
		if currMax > maxSofar{
			maxSofar = currMax
		}
	}
	return maxSofar
}

func main() {
	array := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	println(findMaxSum(array))
}
