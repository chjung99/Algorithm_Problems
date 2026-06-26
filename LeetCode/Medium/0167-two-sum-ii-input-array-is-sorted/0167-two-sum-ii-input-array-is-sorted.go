func twoSum(numbers []int, target int) []int {

    n := len(numbers)
    left := 0
    right := n - 1

    for (left != right) {
        sum := numbers[left] + numbers[right]
        if (sum == target) {
            break
        } else if (sum < target) {
            left += 1
        } else {
            right -= 1
        }
    }
    return []int{left + 1, right + 1}
}