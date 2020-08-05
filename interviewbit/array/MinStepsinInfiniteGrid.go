package array

func GetMinSteps(x []int64, y []int64) int64 {
	var min int64
	for i := 0; i < len(x)-1; i++ {
		xCord := abs(x[i+1] - x[i])
		yCord := abs(y[i+1] - y[i])
		min += max(xCord, yCord)
	}
	return min
}

func max(m int64, n int64) int64 {
	if m > n {
		return m
	}
	return n
}

func abs(n int64) int64 {
	if n > 0 {
		return n
	}
	return -n
}
