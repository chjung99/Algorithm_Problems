func getSneakyNumbers(nums []int) []int {
    snks := make([]int, 0)
    appearance := make(map[int]bool)
    
    for _, num := range nums {
        _, ok := appearance[num]

        if (!ok) {
            appearance[num] = true
        } else {
            snks = append(snks, num)
        }
    }
    
    return snks
}