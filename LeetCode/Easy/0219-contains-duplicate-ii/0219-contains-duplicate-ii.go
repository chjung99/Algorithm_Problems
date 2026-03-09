func containsNearbyDuplicate(nums []int, k int) bool {
    n := len(nums)
    d := min(n-1, k)

    m := make(map[int]int)

    for i := 0; i <= d; i++ {
        val, exists := m[nums[i]]
        if (exists && val == 1) {
            return true
        } else {
            m[nums[i]] = val + 1
        }
    }

    for i := 1; i + d < n; i++ {
        m[nums[i-1]] -= 1
        val, exists := m[nums[i+d]]
        if (exists && val == 1) {
            return true
        } else {
            m[nums[i+d]] = val + 1
        }
    }
    return false
}