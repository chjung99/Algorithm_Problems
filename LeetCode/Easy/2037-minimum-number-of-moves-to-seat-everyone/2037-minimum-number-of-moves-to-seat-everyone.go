func minMovesToSeat(seats []int, students []int) int {
    cnt := 0
    n := len(students)

    slices.Sort(students)
    slices.Sort(seats)

    for i := 0; i < n; i++ {
        cnt += max(students[i]-seats[i], seats[i]-students[i])
    }
    return cnt
}