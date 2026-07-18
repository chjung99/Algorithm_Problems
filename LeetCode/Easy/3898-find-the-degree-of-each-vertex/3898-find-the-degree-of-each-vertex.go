func findDegrees(matrix [][]int) []int {
    n := len(matrix)
    deg := make([]int, n)

    for i := 0; i < n; i++ {
        cnt := 0
        for j := 0; j < n; j++ {
            if (matrix[i][j] == 1) {
                cnt += 1
            }
        }
        deg[i] = cnt
    }
    return deg
}