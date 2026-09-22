func findThePrefixCommonArray(A []int, B []int) []int {
    arr := make([]int, 0)
    n := len(A)

    for i := 0; i < n; i++ {
        cnt := 0
        for j := 0; j < i; j++ {
            if (B[j] == A[i]) {
                cnt += 1
            }
        }

        for k := 0; k <= i; k++ {
            if (B[i] == A[k]) {
                cnt += 1
            }
        }

        if i == 0 {
            arr = append(arr, cnt)
        } else {
            arr = append(arr, arr[i-1] + cnt)
        }
    }
    return arr
}