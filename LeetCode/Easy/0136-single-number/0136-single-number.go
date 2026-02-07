func singleNumber(nums []int) int {
    m := make(map[int]int)

    for _, key := range(nums) {
        val, exists := m[key]

        if (exists) {
            m[key] = val + 1
        } else {
            m[key] = 1
        }
    }

    for key, val := range(m) {
        if (val == 1) {
            return key
        }
    }
    return -1
}