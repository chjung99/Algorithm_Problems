func countGoodTriplets(arr []int, a int, b int, c int) int {
    cnt := 0
    n := len(arr)

    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            for k := j+1; k < n; k++ {
                if max(arr[i] - arr[j], arr[j] - arr[i]) <= a && max(arr[j] - arr[k], arr[k] - arr[j]) <= b && max(arr[i] - arr[k], arr[k] - arr[i]) <= c {
                    cnt ++
                }
            }
        }
    }
    return cnt
}