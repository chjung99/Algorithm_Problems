func isUgly(n int) bool {
    if (n <= 0) {
        return false
    }
    for i := 2; (i*i) <= n; i++ {
        for (n % i == 0) {
            if (i != 2 && i != 3 && i != 5) {
                return false
            }
            n = n / i
        }
    }
    if (n > 1) {
        if (n != 2 && n != 3 && n != 5) {
            return false
        }
    }
    return true
}