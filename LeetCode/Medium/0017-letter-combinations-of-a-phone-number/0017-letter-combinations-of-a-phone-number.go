func letterCombinations(digits string) []string {
    ret := make([]string, 0)
    m := make([][]byte, 10)
    m[0] = make([]byte, 0)
    m[1] = make([]byte, 0)
    m[2] = []byte{'a', 'b', 'c'}
    m[3] = []byte{'d', 'e', 'f'}
    m[4] = []byte{'g', 'h', 'i'}
    m[5] = []byte{'j', 'k', 'l'}
    m[6] = []byte{'m', 'n', 'o'}
    m[7] = []byte{'p', 'q', 'r', 's'}
    m[8] = []byte{'t', 'u', 'v'}
    m[9] = []byte{'w', 'x', 'y', 'z'}

    dfs(digits, m, &ret, 0, len(digits), "")

    return ret
}

func dfs(digits string, m [][]byte, ret *[]string, depth int, n int, out string){
    if (depth == n) {
        // fmt.Println(out)
        *ret = append(*ret, out)
        return
    }
    idx := int(digits[depth] - '0')
    for _, char := range m[idx] {
        dfs(digits, m, ret, depth + 1, n, out + string(char))
    }
}