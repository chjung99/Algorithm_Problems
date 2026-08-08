type Point struct {
    X int
    Y int
} 

func convert(s string, numRows int) string {
    m := make(map[Point]byte)

    x := 0
    y := 0

    if (numRows == 1) {
        return s
    }

    for i := 0; i < len(s); i++ {
        // fmt.Println(x, y, string(s[i]), (i / (numRows-1)) % 2 == 0)
        m[Point{x, y}] = s[i]

        if ((i / (numRows-1)) % 2 == 0) {
            x += 1
        } else {
            x -= 1
            y += 1
        }
        // 1, 2, 3,
        // 4, 5, 6,
        // 7, 8, 9,
        // 10, 11, 12
        // 13,
    }
    ret := make([]byte, 0)
    for i := 0; i < numRows; i++ {
        for j := 0; j <= y; j++ {
            val, exist := m[Point{i, j}]
            if (exist) {
                ret = append(ret, val)
            }
        }
    }
    return string(ret)
}
