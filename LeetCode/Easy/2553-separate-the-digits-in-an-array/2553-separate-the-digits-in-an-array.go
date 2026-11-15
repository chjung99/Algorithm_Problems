func separateDigits(nums []int) []int {
    digits := make([]int, 0)

    for _, num := range nums {
        tmp := make([]int, 0)
        for num != 0 {
            tmp = append(tmp, num % 10)
            num /= 10
        }
        for i := len(tmp) - 1; i >= 0; i-- {
            digits = append(digits, tmp[i])
        }
    }
    return digits
}