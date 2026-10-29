func decompressRLElist(nums []int) []int {
    list := make([]int, 0)

    for i := 0; 2*i + 1 < len(nums); i++ {
        for j := 0; j < nums[2*i]; j ++ {
            list = append(list, nums[2*i+1])
        }
    }
    return list
}