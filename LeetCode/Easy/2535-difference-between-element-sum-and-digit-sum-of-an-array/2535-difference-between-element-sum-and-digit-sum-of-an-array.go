func differenceOfSum(nums []int) int {
    eSum := 0
    dSum := 0

    for _, num := range nums {
        eSum += num
        for num != 0 {
            dSum += num % 10
            num /= 10
        }
    }
    return max(dSum - eSum, eSum - dSum)
}