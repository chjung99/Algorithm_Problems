func isUgly(n int) bool {
    if (n <= 0) {
        return false
    }
    for i := 2; (i*i) <= n; i++ {
        for (n % i == 0) {
            fmt.Println(i)
            if (i != 2 && i != 3 && i != 5) {
                fmt.Println(i)
                return false
            }
            n = n / i
        }
    }
    if (n > 1) {
        if (n != 2 && n != 3 && n != 5) {
            fmt.Println(n)
            return false
        }
    }
    return true
}