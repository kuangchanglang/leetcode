func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))
    res := make([]int, 2, 2)
    
    for i, a := range nums{
        if idx, ok := m[target-a]; ok{
            res[0], res[1] = idx, i
            return res
        }
        m[a] = i
    }
    return res
}
