package main

import (
	"fmt"
)

type sumT func(int, int) int

func sum(a int, b int) int {
	return a + b
}

func total(t sumT, p1 int, p2 int) int {
	return t(p1, p2)
}

func ts(s *[]int) {
	*s = append(*s, 4, 5, 6)
}

func main() {
	s := []int{1, 2, 3}

	fmt.Println(a, s)
}
