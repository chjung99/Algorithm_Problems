func findPermutationDifference(s string, t string) int {
    diff := 0

    sIdx := make(map[rune]int)
    tIdx := make(map[rune]int)

    for i, c := range s {
        sIdx[c] = i
    }

    for i, c := range t {
        tIdx[c] = i
    }
    

    for _, c := range s {
        diff += max(sIdx[c] - tIdx[c], tIdx[c] - sIdx[c])
    }
    return diff
}