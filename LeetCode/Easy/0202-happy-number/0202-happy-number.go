func isHappy(n int) bool {
    m := make(map[int]bool)
    
    for n != 1 {
        exist, val := m[n]
        if (exist && val) {
            return false
        }
        m[n] = true
        x := 0
        for n != 0 {
            x += (n % 10) * (n % 10)
            n = n / 10
        }
        n = x
    }
    return true
}