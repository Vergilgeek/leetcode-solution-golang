package main
/**
公共方法
 */
import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 9}
	var target = 9
	fmt.Println(binarySearch(nums, target))
}
/**
二分查找
 */
func binarySearch(nums []int, target int) int {
	var low = 0
	var high = len(nums) - 1
	for low <= high {
		var mid = low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		}
	}
	return -1
}
