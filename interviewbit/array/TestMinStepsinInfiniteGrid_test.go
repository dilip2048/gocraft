package array

import (
	"testing"
)

func TestMinStepsInInfiniteGrid(t *testing.T) {
	xCoordinate := []int64{4, 1, 4, 10}
	yCoordinate := []int64{6, 2, 5, 12}
	steps := GetMinSteps(xCoordinate, yCoordinate)
	if steps != 2 {
		t.Errorf("Expected: 2, Got: %v", steps)
	}
}
