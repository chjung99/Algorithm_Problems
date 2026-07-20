func solve(board [][]byte)  {
    m := len(board)
    n := len(board[0])
    visit := make([][]bool, m)
    for i := 0; i < m; i++ {
        visit[i] = make([]bool, n)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++{
            if (board[i][j] == 'O' && !visit[i][j]) {
                bfs(i, j, &board, &visit, m, n)
            }
        }
    }
}

func bfs(x int, y int, board *[][]byte, visit *[][]bool, m int, n int) {
    dx := [4]int{-1, 0, 1, 0}
    dy := [4]int{0, 1, 0, -1}

    isEdge := false
    st := make([][2]int, 0)
    st = append(st, [2]int{x, y})
    q := make([][2]int, 0)
    q = append(q, [2]int{x, y})

    (*visit)[x][y] = true

    isEdge = (x == 0 || x == m-1) || (y == 0 || y == n-1)

    for (len(st) != 0) {
        top := st[len(st)-1]
        st = st[:len(st)-1]

        cx, cy := top[0], top[1]
        for i := 0; i < 4; i++ {
            nx, ny := cx + dx[i], cy + dy[i]
            if (nx < 0 || nx >= m || ny < 0 || ny >= n || (*visit)[nx][ny] || (*board)[nx][ny] == 'X') {
                continue
            }
            (*visit)[nx][ny] = true
            st = append(st, [2]int{nx, ny})
            q = append(q, [2]int{nx, ny})

            isEdge = isEdge || (nx == 0 || nx == m-1) || (ny == 0 || ny == n-1)
        }
    }
    if (!isEdge) {
        for (len(q) != 0) {
            top := q[len(q)-1]
            q = q[:len(q)-1]

            cx, cy := top[0], top[1]
            (*board)[cx][cy] = 'X'
        }
    }
}