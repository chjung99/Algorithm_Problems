func subsetXORSum(nums []int) int {
    sum := 0
    n := len(nums)
    for i := 1; i <= n; i++ {
        getSubset(&[]int{}, 0, i, nums, &sum, n)
    }

    return sum
}

func getSubset(subset *[]int, idx int, size int, nums []int, sum *int, n int) {
    if (len(*subset) == size) {
        *sum += calcXOR(*subset)
        return
    }

    for i := idx; i < len(nums); i++ {
        (*subset) = append((*subset), nums[i])
        getSubset(subset, i + 1, size, nums, sum, n)
        *subset = (*subset)[:len(*subset)-1]
    }
}

func calcXOR(arr []int) int {
    ret := 0
    for _, num := range arr {
        ret ^= num
    }
    return ret
}