func canConstruct(ransomNote string, magazine string) bool {
    counts := make(map[rune]int)

    for _, char := range magazine {
        counts[char]++
    }

    for _, char := range ransomNote {
        if (counts[char] == 0) {
            return false
        }
        counts[char]--
    }
    return true
}