package main

import "fmt"

// function to add 1 to the number
func addOne(s []int) []int{
	n := len(s)

	//add 1 to last digit and store the carry
	s[n-1] = s[n-1] + 1
	carry := s[n-1] / 10
	s[n-1] = s[n-1] % 10

	//loop from second last element
	for i := n - 2; i >= 0; i-- {
		if carry == 1 {
			s[i] = s[i] + 1
			carry = s[i] / 10
			s[i] = s[i] % 10
		}
	}

	if carry == 1 {
		s1 := []int{1}
		s1 = append(s1, s...)
		return  s1
	}
	return s
}

func main() {
	slice := []int{9, 9, 9}
	s := addOne(slice)
	print(fmt.Sprintf("%v", s))
}
