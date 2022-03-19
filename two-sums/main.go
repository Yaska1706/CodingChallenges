package main

import (
	"fmt"
	"sort"
)

func noobtwoSum(nums []int, target int) []int {
	sum := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {

				sum[0] = nums[i]
				sum[1] = nums[j]
			}
		}
	}
	return sum

}

func trytwosum(nums []int, result int) []int {
	sum := make([]int, 0)
	min := 0
	max := len(nums) - 1
	nums = sort.IntSlice(nums)

	fmt.Print("starting...")
	for min < max {
		if nums[min]+nums[max] == result {
			sum = append(sum, nums[min], nums[max])
			return sum
		} else if nums[min]+nums[max] < result {
			min = +1

		} else {
			max = max - 1
		}
	}
	return sum
}

func hashtwosum(nums []int, result int) []int {
	hashmap := map[int]int{}
	sum := make([]int, 0)

	for i, num := range nums {
		potentialmatch := result - num
		if j, found := hashmap[potentialmatch]; found {
			sum = append(sum, nums[i], nums[j])
			return sum
		}
		hashmap[num] = i
	}
	return sum
}
