func differenceOfSums(n int, m int) int {
    num1 := 0
    num2 := 0

    for i := 1; i <= n; i++ {
        if (i % m != 0) {
            num1 += i
        } else {
            num2 += i
        }
    }
    return num1 - num2
}