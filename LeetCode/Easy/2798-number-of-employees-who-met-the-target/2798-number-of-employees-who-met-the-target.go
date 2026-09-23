func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    cnt := 0

    for _, hour := range hours {
        if (hour >= target) {
            cnt += 1
        }
    }
    return cnt
}