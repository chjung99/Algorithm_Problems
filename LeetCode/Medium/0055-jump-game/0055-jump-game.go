func canJump(nums []int) bool {
    n := len(nums)
    visit := make([]bool, n)
    var stack []int

    stack = append(stack, 0)
    visit[0] = true

    for len(stack) != 0 {
        top := len(stack) - 1
        cur := stack[top]
        stack = stack[:top]
        
        if (cur >= n) {
            continue
        }

        maxLength := nums[cur]

        for i := 1; i <= maxLength; i++ {
            if (cur + i >= n || visit[cur + i]) {
                continue
            }
            visit[cur + i] = true
            stack = append(stack, cur + i)
        }
    }
    return visit[n-1]
}