func countPoints(points [][]int, queries [][]int) []int {
    arr := make([]int, 0)

    for _, query := range queries {
        x, y, r := query[0], query[1], query[2]
        cnt := 0
        for _, point := range points {
            xi, yi := point[0], point[1]
            if (x - xi) * (x - xi) + (y - yi) * (y - yi) <= r * r {
                cnt += 1 
            }
        }
        arr = append(arr, cnt)
    }
    return arr
}