func countDigits(num int) int {
    cnt := 0
    x := num

    for x != 0 {
        if num % (x % 10) == 0 {
            cnt ++
        }
        x = x / 10
    }
    return cnt
}