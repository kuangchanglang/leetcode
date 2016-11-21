func countPrimes(n int) int {
    cnt := 0
    s := make([]bool, n, n)
    for i := 2; i < n; i++ {
        if s[i] {
            continue
        }
        for j := 2; i*j< n;j++{
            s[i*j] = true
        }
    }
    
    for i := 2; i < n; i++{
        if !s[i] {
            cnt++
        }
    }
    return cnt        
}
