package main

import "fmt"

func quickSort(s []int)[]int{
	if len(s) <= 1 {
		return s
	}
	pivot := s[0]
	i, j := 1, len(s)-1
	for i<=j {
		if s[i] > pivot && s[j] < pivot{
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}else if s[i] < pivot{
			i++
		}else{
			j--
		}
	}
	s[0], s[j] = s[j], s[0]
	quickSort(s[0:i])
	quickSort(s[i:])
	return s
}

func insertSort(s []int)[]int{
	for i, _ := range s{
		for j:=0; j<i; j++{
			if s[j]>s[i]{
				// insertSort it here
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}

func bubbleSort(s []int)[]int{
	for i, _ := range s{
		for j:=i+1; j < len(s);j++{
			if s[i] > s[j]{
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}

func selectSort(s []int) []int{
	for i, _ := range s{
		min_idx := i
		for j:=i+1; j<len(s);j++{
			if s[j]<s[min_idx]{
				min_idx = j
			}
		}
		s[i], s[min_idx] = s[min_idx], s[i]
	}
	return s
}

func mergeSort(s []int)[]int{
	if len(s) < 2{
		return s
	}
	a := mergeSort(s[0:len(s)/2])
	b := mergeSort(s[len(s)/2:])
	return merge(a, b)
}

func merge(a, b []int)[]int{
	res := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for ; i<len(a) && j<len(b);{
		if a[i]<b[j]{
			res = append(res, a[i])
			i++
		}else{
			res = append(res, b[j])
			j++
		}
	}
	for ;i<len(a);i++{
		res = append(res, a[i])
	}

	for ;j<len(b);j++{
		res = append(res, b[j])
	}
	return res
}


func main() {
	fmt.Println(bubbleSort([]int{8, 5, 9, 1, 2, 6, 3, 4, 7}))
	fmt.Println(insertSort([]int{8, 5, 9, 1, 2, 6, 3, 4, 7}))
	fmt.Println(selectSort([]int{8, 5, 9, 1, 2, 6, 3, 4, 7}))
	fmt.Println(mergeSort([]int{8, 5, 9, 1, 2, 6, 3, 4, 7}))
	fmt.Println(quickSort([]int{8, 5, 9, 1, 2, 6, 3, 4, 7}))
	fmt.Println("vim-go")
}
