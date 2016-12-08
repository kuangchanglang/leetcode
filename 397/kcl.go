package main

import "fmt"

func integerReplacement(n int) int {
	m := make(map[int]int)
	return cnt(n, m)
}

func cnt(n int, m map[int]int) int {
	if re, ok := m[n]; ok{
		return re
	}
	if n == 1{
		return 0;
	}
	if n % 2 == 0{
		m[n/2] = cnt(n/2, m)
		return m[n/2]+ 1
	}
	m[n-1] = cnt(n-1, m)
	m[n+1] = cnt(n+1, m)
	if m[n-1] < m[n+1]{
		return m[n-1]+1
	}
	return m[n+1]+1
}

func main() {
	fmt.Println(integerReplacement(1))
	fmt.Println(integerReplacement(7))
	fmt.Println(integerReplacement(8))
}
