func hammingDistance(x int, y int) int {
    return countSet(x^y)
}

func countSet(x int) int{
    cnt := 0
    for x > 0 {
        if (x % 2 == 1) {
            cnt += 1
        }
        x = x / 2
    }
    return cnt
}

