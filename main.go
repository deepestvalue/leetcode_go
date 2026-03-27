package main

import (
	"fmt"
	p125 "leetcode_go/problems/problem125validpalindrome"
	p1 "leetcode_go/problems/problem1twosum"
	p344 "leetcode_go/problems/problem344reversestring"
	p3 "leetcode_go/problems/problem3longestsubstring"
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

	// Problem 125
	fmt.Println("\nProblem 125 Valid Palindrome")
	fmt.Println(p125.IsPalindrome("hi"))
	fmt.Println(p125.IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(p125.IsPalindrome("race a car"))
	fmt.Println(p125.IsPalindrome(" "))

	// Problem 344
	fmt.Println("\nProblem 344 Reverse String")
	s := []byte("hello")
	p344.ReverseString(s)
	fmt.Println(string(s))

	s = []byte("what's up people")
	p344.ReverseString(s)
	fmt.Println(string(s))

	s = []byte("abcd")
	p344.ReverseString(s)
	fmt.Println(string(s))

	// Problem 3
	fmt.Println("\n Problem 3 Longest Substring Without Repeating Characters")
	fmt.Println(p3.LengthOfLongestSubstring("hi"))
	fmt.Println(p3.LengthOfLongestSubstring("abcabcbb"))
	fmt.Println(p3.LengthOfLongestSubstring("bbbbb"))
	fmt.Println(p3.LengthOfLongestSubstring("pwwkew"))
	fmt.Println(p3.LengthOfLongestSubstring("dvdf"))
	fmt.Println(p3.LengthOfLongestSubstring("tmmzuxt"))
}
