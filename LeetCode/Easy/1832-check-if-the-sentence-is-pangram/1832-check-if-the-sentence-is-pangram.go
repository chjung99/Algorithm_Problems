func checkIfPangram(sentence string) bool {
    chk := make(map[rune]bool)
    for _, char := range sentence {
        chk[char] = true
    }
    return len(chk) == 26
}