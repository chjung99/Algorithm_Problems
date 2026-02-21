func generate(numRows int) [][]int {
    arr := make([][]int, numRows)
    for i := 0; i < numRows; i++ {
        arr[i] = make([]int, i+1)
        for j := 0; j < i+1; j++ {
            if (j == 0 || j == i) {
                arr[i][j] = 1
            } else {
                arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
            }
        }
    }
    return arr
}