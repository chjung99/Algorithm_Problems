func findTheDifference(s string, t string) byte {
    sm := make(map[rune]int)
    tm := make(map[rune]int)

    for _, char := range s {
        _, exists := sm[char]
        if (!exists) {
            sm[char] = 0
        }
        sm[char] += 1
    }

    for _, char := range t {
        _, exists := tm[char]
        if (!exists) {
            tm[char] = 0
        }
        tm[char] += 1
    }

    diff := t[0]

    for _, char := range t {
        sCnt, sExists := sm[char]
        tCnt, _ := tm[char]
        
        if (!sExists || sCnt != tCnt) {
            diff = byte(char)
            break
        }
    }
    return diff
}