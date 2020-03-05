package main

import "log"

func twoSum(nums []int, target int) interface{} {
	m := make(map[int]int)
	for index, value := range nums {
		_, ok := m[target - value]
		if ok {
			return true
		}
		m[value] = index
	}
	return false
}

func main() {
	array := []int{1, 4, 45, 6, 10, 8}
	sum := 16
	exist := twoSum(array, sum)
	log.Println(exist)
}
