func sumOfTheDigitsOfHarshadNumber(x int) int {
    h := 0
    num := x

    for num != 0 {
        h += num % 10
        num /= 10
    }

    if x % h == 0 {
        return h 
    }
    return -1
}