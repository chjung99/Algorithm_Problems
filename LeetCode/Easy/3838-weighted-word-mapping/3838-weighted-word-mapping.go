func mapWordWeights(words []string, weights []int) string {
    ret := make([]rune, 0)
    for _, word := range words {
        weight := 0
        for _, ch := range word {
            weight += weights[int(ch - 'a')]
        }

        ret = append(ret, rune('a' +(25 - weight % 26)))
    }
    
    return string(ret)
}