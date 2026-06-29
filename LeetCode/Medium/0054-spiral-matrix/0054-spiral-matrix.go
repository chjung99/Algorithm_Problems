func spiralOrder(matrix [][]int) []int {
    n := len(matrix)
    m := len(matrix[0])
    ans := make([]int, 0)
    visit := make([][]bool, n)
    dx := []int{0, 1, 0, -1}
    dy := []int{1, 0, -1, 0}

    for i := 0; i < n; i++ {
        visit[i] = make([]bool, m)
    }

    x := 0
    y := 0
    d := 0
    cnt := 0
    nx := 0
    ny := 0
    

    for (cnt != n * m) {
        ans = append(ans, matrix[x][y])
        visit[x][y] = true
        cnt += 1

        nx = x + dx[d]
        ny = y + dy[d]

        if (nx < 0 || nx >= n || ny < 0 || ny >= m || visit[nx][ny]) {
            d = (d+1) % 4
            nx = x + dx[d]
            ny = y + dy[d]
        }

        x = nx
        y = ny
    }

    

    return ans
}

