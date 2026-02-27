func intersection(nums1 []int, nums2 []int) []int {
    num := make(map[int]bool)
    visit := make(map[int]bool)
    var ret []int

    for _, val := range nums1 {
        num[val] = true
    }

    for _, val := range nums2 {
        if (num[val] && !visit[val]) {
            ret = append(ret, val)
            visit[val] = true
        }
    }
    return ret

}