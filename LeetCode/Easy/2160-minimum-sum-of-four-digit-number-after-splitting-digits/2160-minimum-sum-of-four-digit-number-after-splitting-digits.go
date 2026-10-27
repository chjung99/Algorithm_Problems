func minimumSum(num int) int {
    digits := make([]int, 0)
    for num != 0 {
        digits = append(digits, num % 10)
        num = num / 10
    }

    slices.Sort(digits)
    x := 0
    y := 0

    for i := 0; i < len(digits); i++ {
        if i % 2 == 0 {
            x *= 10
            x += digits[i]
        } else {
            y *= 10
            y += digits[i]
        }
    }
    return x + y
}