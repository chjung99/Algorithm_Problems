func mostWordsFound(sentences []string) int {
    maxNum := 0
    for _, s := range sentences {
        maxNum = max(maxNum, len(strings.Split(s, " ")))
    }
    return maxNum
}