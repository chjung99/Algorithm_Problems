func numOfStrings(patterns []string, word string) int {
    res := 0
    for _, pattern := range patterns {
        if strings.Contains(word, pattern) {
            res += 1
        }
    }
    return res
}