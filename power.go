package main

func power(x int64, y int64) int64 {
	if y == 0 {
		return 1
	}
	if y%2 == 0 {
		return power(x, y/2) * power(x, y/2)
	} else {
		return x * power(x, y/2) * power(x, y/2)
	}
}

func main() {
	result := power(2, 5)
	println(result)
}
