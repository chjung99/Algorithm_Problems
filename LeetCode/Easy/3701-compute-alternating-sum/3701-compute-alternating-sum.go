func alternatingSum(nums []int) int {
    sum := 0

    for idx, num := range nums {
        if (idx % 2 == 0) {
            sum += num
        } else {
            sum -= num
        }
    }

    return sum
}