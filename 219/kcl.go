package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i:=0;i<len(nums);i++{
		if j, ok := m[nums[i]]; ok && i-j<=k{
			return true
		}
		m[nums[i]]=i
	}
	return false
}
func main() {
	fmt.Println("vim-go")
}
