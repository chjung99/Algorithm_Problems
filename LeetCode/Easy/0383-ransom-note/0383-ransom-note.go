func canConstruct(ransomNote string, magazine string) bool {
    m := make(map[rune]int)

    for _, char := range magazine {
        _, exists := m[char]
        if (!exists) {
            m[char] = 1
        } else {
            m[char] += 1
        }
    }

    for _, char := range ransomNote {
        val, exists := m[char]
        if (!exists || val == 0) {
            return false
        } else {
            m[char] -= 1
        }
    }


    return true
}