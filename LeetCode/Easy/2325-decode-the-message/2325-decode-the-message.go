func decodeMessage(key string, message string) string {
    table := make(map[rune]rune)
    cnt := 0
    ret := ""
    for _, char := range key {
        if char == ' ' {
            continue
        }
        if _, exists := table[char]; exists {
            continue
        }
        table[char] = rune('a' + cnt)
        cnt += 1
    }
    for _, char := range message {
        if char == ' ' {
            ret += string(char)
            continue
        }
        ret += string(table[char])
    }
    return ret
}