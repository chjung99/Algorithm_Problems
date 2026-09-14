func countConsistentStrings(allowed string, words []string) int {
    m := make(map[rune]bool)
    cnt := 0

    for _, char := range allowed {
        if _, exists := m[char]; !exists {
            m[char] = true
        }
    }

    for _, word := range words {
        isConsistent := true
        for _, char:= range word {
            if _, exists := m[char]; !exists {
                isConsistent = false
                break
            }
        }
        if (isConsistent) {
            cnt += 1
        }
    }

    return cnt
}