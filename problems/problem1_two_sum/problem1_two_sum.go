package problems1_two_sum

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func TwoSumSol2(nums []int, target int) []int {
	wanted := make(map[int]int)
	for i := range nums {
		wanted[nums[i]] = i
	}

	for i := range nums {
		complement := target - nums[i]

		if j, ok := wanted[complement]; ok && j != i {
			return []int{i, j}
		}
	}

	return []int{}
}

func TwoSumSol3(nums []int, target int) []int {
	wanted := make(map[int]int)
	for i := range nums {
		complement := target - nums[i]

		if j, ok := wanted[complement]; ok {
			return []int{i, j}
		}

		wanted[nums[i]] = i
	}

	return []int{}
}
