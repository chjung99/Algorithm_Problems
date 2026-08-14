func isStrictlyPalindromic(n int) bool {
    ret := false

    b := n - 2

    for (b >= 2) {
        arr := make([]int, 0)
        x := n
        for (x != 0) {
            arr = append(arr, x % b)
            x = x / b
        }
        b -= 1

        for i := 0 ; i < len(arr) / 2; i++ {
            if (arr[i] != arr[len(arr)-1-i]) {
                ret = false
                break
            }
        }
    }

    return ret
}