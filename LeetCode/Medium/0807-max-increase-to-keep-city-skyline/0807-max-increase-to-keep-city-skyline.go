func maxIncreaseKeepingSkyline(grid [][]int) int {
    n := len(grid)
    inc := 0
    maxRow := func(r int) int {
        m := grid[r][0]
        for i := 0; i < n; i++ {
            m = max(m, grid[r][i])
        }
        return m
    }

    maxCol := func(c int) int {
        m := grid[0][c]
        for i := 0; i < n; i++ {
            m = max(m, grid[i][c])
        }
        return m
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            mr := maxRow(i)
            mc := maxCol(j)

            inc += min(mr, mc) - grid[i][j]
        }
    }
    return inc
}