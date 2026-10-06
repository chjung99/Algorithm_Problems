func numberOfMatches(n int) int {
    num := 0

    for n >= 2 {
        if n % 2 == 0 {
            num += n / 2
            n = n / 2
        } else {
            num += (n - 1) / 2
            n = (n - 1) / 2 + 1
        }
    }
    return num
}