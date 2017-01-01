package main

import "fmt"

func longestPalindrome(s string) int {
	l := 0
	m := make(map[rune]int, 0)
	for _, c := range s{
		if m[c] > 0{
			m[c] = 0
			l += 2
		}else{
			m[c] = 1
		}
	}

	if l < len(s){
		l += 1
	}
	return l
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}
