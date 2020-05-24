package main

import "math"

// This method with return maximum absolute difference
func getMaxAbsoluteDifference(A []int) int{
	max := 0
	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < len(A); j++ {
			sum := math.Abs(float64(A[i]-A[j])) + math.Abs(float64(i-j))
			if int(sum) > max {
				max = int(sum)
			}
		}
	}
	return max
}

func main() {
	array := []int{ -70, -64, -6, -56, 64,
		61, -57, 16, 48, -98 }
	max := getMaxAbsoluteDifference(array)
	println(max)
}
