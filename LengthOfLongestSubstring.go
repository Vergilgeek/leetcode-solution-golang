package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk + 1 < n && m[s[rk+1]] == 0 {
			m[s[rk + 1]]++
			rk ++
		}
		ans = int(math.Max(float64(ans), float64(rk-i+1)))
	}
	return ans
}
