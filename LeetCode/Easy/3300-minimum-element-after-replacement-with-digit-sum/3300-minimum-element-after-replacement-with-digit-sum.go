func minElement(nums []int) int {
    rep := make([]int, 0)
    minE := 0

    for _, num := range nums {
        digitSum := getDigitSum(num)
        if (minE == 0) {
            minE = digitSum
        } else {
            minE = min(minE, digitSum)
        }
        rep = append(rep, digitSum)
    }

    return minE
}

func getDigitSum(digit int) int {
    sum := 0
    for (digit > 0) {
        sum += digit % 10
        digit = digit / 10
    }
    return sum
}