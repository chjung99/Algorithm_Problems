func transformArray(nums []int) []int {
    arr := make([]int, 0)

    for _, num := range nums {
        if (num % 2 == 0) {
            arr = append(arr, 0)
        } else {
            arr = append(arr, 1)
        }
    }
    sort.Ints(arr)
    return arr
}