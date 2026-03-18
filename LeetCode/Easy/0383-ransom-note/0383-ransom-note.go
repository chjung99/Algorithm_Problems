func canConstruct(ransomNote string, magazine string) bool {
    x := make(map[rune]int)
    y := make(map[rune]int)

    for _, val := range(ransomNote) {
        cnt, exists :=  x[val]

        if (!exists) {
            x[val] = 0
        } else {
            x[val] = cnt + 1
        }
    }

    for _, val := range(magazine) {
        cnt, exists :=  y[val]

        if (!exists) {
            y[val] = 0
        } else {
            y[val] = cnt + 1
        }
    }

    for key, cntX := range x {
        cntY, exists := y[key]
        if (!exists) {
            return false
        } else if (cntY < cntX){
            return false
        }
    }
    return true
}