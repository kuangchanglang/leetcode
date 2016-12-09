package main

import "fmt"

func circularArrayLoop(nums []int) bool {
	len := len(nums)
	for i, n := range nums{
		nums[i] = ((n+i)%len+len)%len
	}

	for i, _ := range nums{
		j:=i
		last, llast := i, -1
		for{
			if nums[j]<0{
				if nums[j]==-1*(i+1) && j != last && j!=llast{
					return true
				}
				break
			}
			next := nums[j]
			nums[j] = -1*(i+1)
			last, llast = j, last
			j = next
		}
	}

	return false
}

func main() {
	fmt.Println(circularArrayLoop([]int{2,-1,1,2,2}))
	fmt.Println(circularArrayLoop([]int{2, -1, 1, -2, -2}))
	fmt.Println(circularArrayLoop([]int{-1,2}))
}
