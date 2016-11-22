
func countBits(num int) []int {
	res := make([]int, num+1, num+1)

	for i := 0;i <= num; i++{
		res[i] = bits(i)
	}

	return res
}

func bits(n int)int{
	cnt := 0
	for n >0{
		cnt += n % 2
		n = n/2
	}
	return cnt
}
