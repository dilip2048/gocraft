package main

/* Function to calculate x raised to the power y */
func power(x int64, n int64) int64 {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		return power(x, n/2) * power(x, n/2)
	} else {
		return x * power(x, n/2) * power(x, n/2)
	}
}

/* Program to test function power */
func main() {
	result := power(2, 5)
	println(result)
}