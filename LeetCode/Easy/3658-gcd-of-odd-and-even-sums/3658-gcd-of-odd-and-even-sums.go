func gcdOfOddEvenSums(n int) int {
    return getGCD(getSumOdd(n), getSumEven(n))
}

func getSumOdd(n int) int {
    ret := 0
    for i := 1; i <= 2 * n - 1; i = i + 2 {
        ret += i
    }
    return ret
}

func getSumEven(n int) int {
    ret := 0
    for i := 2; i <= 2 * n; i = i + 2 {
        ret += i
    }
    return ret
}

func getGCD(x int , y int) int {
    fmt.Println(x, y)
    z := min(x, y)

    for (true) {
        if (x % z == 0 && y % z == 0) {
            break
        }
        z -= 1
    }
    return z
}