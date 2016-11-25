package main
import "fmt"

// TODO: refactoring
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)

	// count p chars
	mp := make(map[byte]int, len(p))
	for i:=0;i<len(p);i++{
		c := p[i]
		if _, ok:=mp[c];!ok {
			mp[c] = 0
		}
		mp[c]++
	}

	m := make(map[byte]int, len(p))
	i, j := 0, 0
	for {
		for {
			if j>=len(s){
				return res
			}
			c := s[j]
			if _, ok := m[c]; !ok{
				m[c] = 0
			}
			m[c]++

			j++
			if mp[c] == 0{ // s[j] not exists in p, clear m and jump i to j
				i=j
				m = make(map[byte]int, len(p))
				break
			}

			if m[c] > mp[c]{ // count char s[c] > p[c], move i to next c
				m[c]--
				for{
					d := s[i]
					i++
					if c == d{
						break
					}
					if _, ok:=m[d]; ok && m[d]>0{
						m[d]--
					}
				}
				break
			}
			if j-i == len(p){ // ok, append result, move i forward
				res = append(res, i)
				m[s[i]]--
				i++
				break
			}
		}
	}
	return res
}

func main(){
	s := "beeaaedcbc"
	p := "c"
	fmt.Println(s)
	fmt.Println(p)
	fmt.Println(findAnagrams(s,p))
}
