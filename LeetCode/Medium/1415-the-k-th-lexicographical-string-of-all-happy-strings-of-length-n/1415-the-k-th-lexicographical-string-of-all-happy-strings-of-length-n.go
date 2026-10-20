func getHappyString(n int, k int) string {
    happyArr := make([]string, 0)
    var getHappyArray func(n int, s string)

    getHappyArray = func(n int, s string) {
        if len(s) == n {
            happyArr = append(happyArr, s)
            return
        }

        for _, char := range []byte{'a', 'b', 'c'} {
            if len(s) >= 1 && s[len(s)-1] == char {
                continue
            }
            getHappyArray(n, s + string(char))
        }
    }
    getHappyArray(n, "")

    if len(happyArr) < k {
        return ""
    }

    return happyArr[k-1]
}