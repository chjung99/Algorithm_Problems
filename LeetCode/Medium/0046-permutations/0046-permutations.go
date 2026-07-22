func permute(nums []int) [][]int {
    ans := make([][]int, 0)
    n := len(nums)
    visit := make([]bool, n)

    for i := 0; i < n; i++ {
        visit[i] = true
        perm(&ans, nums, &visit, 1, &[]int{nums[i]})
        visit[i] = false
    }
    return ans
}

func perm(ans *[][]int, nums []int, visit *[]bool, depth int, p *[]int) {
    if (depth == len(nums)){
        q := make([]int, 0)
        for i := 0; i < len(*p); i++ {
            q = append(q, (*p)[i])
        }
        *ans = append(*ans, q)
        return
    }

    for i := 0; i < len(nums); i++ {
        if ((*visit)[i]) {
            continue
        }
        (*visit)[i] = true
        *p = append(*p, nums[i])
        perm(ans, nums, visit, depth+1, p)
        (*p) = (*p)[:len(*p)-1]
        (*visit)[i] = false
    }
}

