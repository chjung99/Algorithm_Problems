func fib(n int) int {
    var f []int
    
    f = append(f, 0)
    f = append(f, 1)
    
    for i := 2; i <= n; i++ {
        f = append(f, f[i-1] + f[i-2])
    }
    return f[n]
}