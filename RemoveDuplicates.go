package main

import "fmt"

/**
26. 删除排序数组中的重复项
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
*/
func main() {
	var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(solution1(nums))
}

/**
方法：双指针法
数组完成排序后，我们可以放置两个指针 ii 和 jj，其中 ii 是慢指针，而 jj 是快指针。只要 nums[i] = nums[j]nums[i]=nums[j]，我们就增加 jj 以跳过重复项。
*/
func solution1(nums []int) int {
	var l = 0
	for i := range nums {
		if nums[l] != nums[i] {
			nums[l+1] = nums[i]
			l++
		}
	}
	return l + 1
}
