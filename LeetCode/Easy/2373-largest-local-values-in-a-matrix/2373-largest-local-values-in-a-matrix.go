func largestLocal(grid [][]int) [][]int {
    n := len(grid)
    maxLocal := make([][]int, n-2)

    findLargestLocal := func(x int, y int) int {
        ret := grid[x][y]

        for i := x - 1 ; i < x + 2; i++ {
            for j := y - 1; j < y + 2; j++ {
                ret = max(ret, grid[i][j])
            }
        }

        return ret
    }

    for i := range n-2 {
        maxLocal[i] = make([]int, n-2)
    }

    for i := 1; i < n - 1; i++ {
        for j := 1; j < n - 1; j++ {
            maxLocal[i-1][j-1] = findLargestLocal(i, j)
        }
    }

    return maxLocal
}