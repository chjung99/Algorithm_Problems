func majorityElement(nums []int) int {
    slices.Sort(nums)
    prev := nums[0]
    cnt := 0
    ansValue := nums[0]
    ans := 0
    for _, num := range(nums) {
        if (num == prev) {
            cnt++
        } else {
            if (ans < cnt) {
                ans = cnt
                ansValue = prev
            }
            prev = num
            cnt = 1
        }
        // fmt.Println(num, cnt, ans)
    }
    if (ans < cnt) {
        ans = cnt
        ansValue = prev
    }

    return ansValue
}