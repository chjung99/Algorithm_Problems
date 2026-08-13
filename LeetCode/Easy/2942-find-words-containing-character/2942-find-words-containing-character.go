func findWordsContaining(words []string, x byte) []int {
    idxArr := make([]int, 0)
    for idx, word := range words {
        flag := 0
        for _, c := range word {
            if (byte(c) == x) {
                flag = 1
                break
            }
        }
        if (flag == 1) {
            idxArr = append(idxArr, idx)
        }
    }
    return idxArr
}