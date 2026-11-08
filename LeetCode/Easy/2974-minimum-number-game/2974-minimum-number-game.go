func numberGame(nums []int) []int {
    arr := make([]int, 0)
    slices.Sort(nums)

    for i := 0; i < len(nums)/2; i++ {
        alice := nums[2*i]
        bob := nums[2*i+1]

        arr = append(arr, bob)
        arr = append(arr, alice)
    }
    return arr
}