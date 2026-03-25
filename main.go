package main

import (
	"fmt"
	p1 "leetcode_go/problems/problem1twosum"
	p2 "leetcode_go/problems/problem2validpalindrome"
	p3 "leetcode_go/problems/problem3reversestring"
)

func main() {
	// Problem 1
	fmt.Println("\nProblem 1 Two Sum")
	res := p1.TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)

	res = p1.TwoSumSol2([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)

	res = p1.TwoSumSol3([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)

	// Problem 2
	fmt.Println("\nProblem 2 Valid Palindrome")
	fmt.Println(p2.IsPalindrome("hi"))
	fmt.Println(p2.IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(p2.IsPalindrome("race a car"))
	fmt.Println(p2.IsPalindrome(" "))

	// Problem 3
	fmt.Println("\nProblem 3 Reverse String")
	s := []byte("hello")
	p3.ReverseString(s)
	fmt.Println(string(s))

	s = []byte("what's up people")
	p3.ReverseString(s)
	fmt.Println(string(s))

	s = []byte("abcd")
	p3.ReverseString(s)
	fmt.Println(string(s))
}
