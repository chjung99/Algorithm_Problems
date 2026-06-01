func jump(nums []int) int {
    n := len(nums)
    var stack []int
    dist := make([]int, n)

    for i := 0; i < n; i++ {
        dist[i] = math.MaxInt
    }

    stack = append(stack, 0)
    dist[0] = 0

    for len(stack) > 0 {
        top := len(stack) - 1
        cur := stack[top]
        stack = stack[:top]
        
        for j := 1; j <= nums[cur]; j++ {
            nxt := cur + j
            if (nxt >= n) {
                continue
            }
            if (dist[nxt] > dist[cur] + 1) {
                dist[nxt] = dist[cur] + 1
                stack = append(stack, nxt)
            }
        }
    }

    return dist[n-1]
}