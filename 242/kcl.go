package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	m := make(map[rune]int, 26)
	for _, c := range s{
		if _, ok := m[c]; !ok{
			m[c] = 0
		}
		m[c]++
	}
	for _, c := range t{
		cnt := m[c]
		if cnt <= 0{
			return false
		}
		m[c]--
	}
	return true
}
func main() {
	fmt.Println("vim-go")
}
