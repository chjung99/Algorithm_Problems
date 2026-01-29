func searchInsert(nums []int, target int) int {
    left := 0
    right := len(nums)

    for left + 1 < right {
        mid := (left + right) / 2
        if (nums[mid] >= target) {
            right = mid
        } else {
            left = mid
        }
    }
    if (left == 0 && nums[left] >= target) {
        return left
    }
    return left + 1
}