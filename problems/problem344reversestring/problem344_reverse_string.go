package problem344reversestring

func ReverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp

		left++
		right--
	}
}
