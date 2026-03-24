package main

import (
	"fmt"
	p1 "leetcode_go/problems/problem1_two_sum"
)

func main() {
	res := p1.TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)

	res = p1.TwoSumSol2([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)

	res = p1.TwoSumSol3([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}
