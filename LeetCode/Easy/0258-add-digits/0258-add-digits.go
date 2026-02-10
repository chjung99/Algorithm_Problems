func addDigits(num int) int {
    ret := 0
    for num > 0 {
        ret += num % 10
        num /= 10
        if (num == 0) {
            // fmt.Println(num, ret)
            if ((ret / 10) == 0) {
                return ret
            } else {
                num = ret
                ret = 0
            }
        }
    }
    return ret
}