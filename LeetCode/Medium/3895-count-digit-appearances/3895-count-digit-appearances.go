func countDigitOccurrences(nums []int, digit int) int {
    cnt := 0
    for _, num := range nums {
        for num != 0 {
            if num % 10 == digit {
                cnt ++
            }
            num /= 10
        }
    }
    return cnt
}