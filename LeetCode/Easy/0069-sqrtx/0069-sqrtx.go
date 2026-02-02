func mySqrt(x int) int {
    left := 0
    right := x

    if (x == 1) {
        return x
    }

    for left + 1 < right {
        mid := int((left + right) / 2)
        fmt.Println(mid)
        if (mid * mid <= x) {
            left = mid
        } else {
            right = mid
        }
    }
    return left
}