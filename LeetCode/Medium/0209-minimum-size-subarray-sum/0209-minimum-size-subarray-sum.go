func minSubArrayLen(target int, nums []int) int {
    // sort.Ints(nums)

    l := 0
    r := 0
    s := nums[l]

    n := len(nums)

    ans := n + 1

    // 1, 2, 2, 3, 3, 4
    fmt.Println(nums)
    for (r != n) {
        if (s == target) {
            fmt.Println(l, r, s)
            ans = min(ans, r - l + 1)
            r += 1
            if (r < n) {
                s += nums[r]
            }
        } else if (s < target) {
            r += 1
            if (r < n) {
                s += nums[r]
            }
        } else {
            fmt.Println(l, r, s)

            ans = min(ans, r - l + 1)
            s -= nums[l]
            l += 1
        }
    }


    if (ans == n + 1) {
        return 0
    }

    return ans
}