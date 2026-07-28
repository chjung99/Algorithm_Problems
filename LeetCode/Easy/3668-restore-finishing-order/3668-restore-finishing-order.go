func recoverOrder(order []int, friends []int) []int {
    m := make(map[int]int)
    arr := make([]int, 0)

    for _, val := range friends {
        m[val] = val
    }

    for _, val := range order {
        _, exists := m[val]

        if (exists){
            arr = append(arr, val)
        }
    }
    return arr
}