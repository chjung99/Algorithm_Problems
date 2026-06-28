func isValidSudoku(board [][]byte) bool {
    for i := 0; i < 9; i++ {
        if (!isValidRow(board, i)){
            return false
        }
    }

    for j := 0; j < 9; j++ {
        if (!isValidCol(board, j)){
            return false
        }
    }

    for i := 0; i <= 6; i = i + 3 {
        for j := 0; j <= 6; j = j + 3 {
            if (!isValidSubBox(board, i, j)){
                return false
            }
        }
    }
    return true
}

func isValidRow(board [][]byte, r int) bool {
    m := make(map[byte]int)

    for i := 0; i < 9; i++ {
        val, exists := m[board[r][i]]
        if (!exists) {
            m[board[r][i]] = 1
        } else {
            m[board[r][i]] = val + 1
        }
    }

    for k, v := range m {
        if (k >= '1' && k <= '9' && v >= 2) {
            return false
        }
    }
    return true
}

func isValidCol(board [][]byte, c int) bool {
    m := make(map[byte]int)

    for i := 0; i < 9; i++ {
        val, exists := m[board[i][c]]
        if (!exists) {
            m[board[i][c]] = 1
        } else {
            m[board[i][c]] = val + 1
        }
    }

    for k, v := range m {
        if (k >= '1' && k <= '9' && v >= 2) {
            return false
        }
    }
    return true
}

func isValidSubBox(board [][]byte, x int, y int) bool {
    m := make(map[byte]int)

    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            val, exists := m[board[x+i][y+j]]
            if (!exists) {
                m[board[x+i][y+j]] = 1
            } else {
                m[board[x+i][y+j]] = val + 1
            }
        }
    }

    for k, v := range m {
        if (k >= '1' && k <= '9' && v >= 2) {
            return false
        }
    }
    return true
}

