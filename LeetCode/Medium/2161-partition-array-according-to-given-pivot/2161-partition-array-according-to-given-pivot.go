func pivotArray(nums []int, pivot int) []int {
    smaller := make([]int, 0)
    equal := make([]int, 0)
    larger := make([]int, 0)

    for _, num := range nums {
        if (num < pivot) {
            smaller = append(smaller, num)
        } else if (num == pivot) {
            equal = append(equal, num)
        } else {
            larger = append(larger, num)
        }
    }
    ans := append(smaller, equal...)
    ans = append(ans, larger...)
    return ans
}