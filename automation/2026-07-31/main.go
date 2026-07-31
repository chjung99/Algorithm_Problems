func numIslands(grid [][]byte) int {
    n := len(grid)
    m := len(grid[0])
    num := 0

    visit := make([][]bool, n)
    for i := 0; i < n; i++ {
        visit[i] = make([]bool, m)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if (!visit[i][j] && grid[i][j] == '1') {
                num += 1
                bfs(i, j, visit, grid, n, m)
            }
        }
    }
    return num
}

func bfs (x int, y int, visit [][]bool, grid [][]byte, n int, m int) {
    dx := []int{0, 1, 0, -1}
    dy := []int{1, 0, -1, 0}

    st := make([][2]int, 0)
    st = append(st, [2]int{x, y})
    visit[x][y] = true

    for (len(st) != 0) {
        cur := st[len(st) - 1]
        st = st[:len(st) - 1]

        for i := 0; i < 4; i++ {
            nx := cur[0] + dx[i]
            ny := cur[1] + dy[i]

            if (nx < 0 || nx >= n || ny < 0 || ny >= m || visit[nx][ny] || grid[nx][ny] == '0') {
                continue
            }

            visit[nx][ny] = true
            st = append(st, [2]int{nx, ny})
        }
    }
}