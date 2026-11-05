func numberOfSteps(num int) int {
    cnt := 0

    for num != 0 {
        if num % 2 == 0 {
            num /= 2
        } else {
            num --
        }
        cnt ++
    }
    return cnt
}