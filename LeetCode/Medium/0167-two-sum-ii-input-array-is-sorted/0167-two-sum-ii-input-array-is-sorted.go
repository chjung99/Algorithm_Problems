func twoSum(numbers []int, target int) []int {
    idx0 := 0
    idx1 := 1
    n := len(numbers)

    for (idx0 != idx1) {
        sum := numbers[idx0] + numbers[idx1]
        if (sum == target) {
            break
        } else if (sum < target) {
            if (idx1 + 1 < n && numbers[idx0] + numbers[idx1 + 1] <= target) {
                idx1 += 1
            } else {
                idx0 += 1
            }
        } else {
            idx1 -= 1
        }
    }
    return []int{idx0 + 1, idx1 + 1}
}