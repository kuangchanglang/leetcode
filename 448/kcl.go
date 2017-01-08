package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	for i, n := range nums {
		for nums[i] != nums[n-1]{
			nums[i] = nums[n-1]
			nums[n-1] = n
			n = nums[i]
		}
	}

	fmt.Println(nums)
	res := make([]int, 0)
	for i, n := range nums {
		if n != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4,3,2,7,8,2,3,1}))
	fmt.Println(findDisappearedNumbers([]int{1,2,3,5,5}))
}
