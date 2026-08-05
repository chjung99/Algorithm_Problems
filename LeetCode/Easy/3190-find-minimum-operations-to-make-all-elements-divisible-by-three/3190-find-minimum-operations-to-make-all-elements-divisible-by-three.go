func minimumOperations(nums []int) int {
    cnt := 0
    for _, val := range nums {
        tar := findNear(val)
        cnt += min(val - tar, tar + 3 - val)
    }
    return cnt
}

func findNear(val int) int {
    for i := val; i >= 0; i-- {
        if (i % 3 == 0) {
            return i
        }
    }
    return 0
}