func subtractProductAndSum(n int) int {
    digits := make([]int, 0)
    for n != 0 {
        digits = append(digits, n % 10)
        n = n / 10
    }

    product := 1
    sum := 0

    for _, digit := range digits {
        product *= digit
        sum += digit
    }

    return product - sum
}