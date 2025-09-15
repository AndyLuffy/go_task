package main

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数

func twoSum(nums []int, target int) []int {
	arr := make([]int, 2)

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				arr = []int{i, j}
			}
		}
	}

	return arr
}
