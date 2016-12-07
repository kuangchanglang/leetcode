package main
import (
	"fmt"
)

func isHappy(n int) bool {
	m := make(map[int]bool)
	for !m[n] {
		m[n] = true
		n = next(n)
		if n == 1{
			return true
		}
	}
	return false
}

func next(n int) int{
	sum := 0
	for n>0 {
		sum += (n%10)*(n%10)
		n /= 10
	}
	return sum
}

func main(){
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(1))
	fmt.Println(isHappy(0))
}
