func maxFreqSum(s string) int {
    maxVowelFreq := 0
    maxConsonantFreq := 0

    freq := make(map[rune]int)

    for _, c := range s {
        _, exists := freq[c]
        if (!exists) {
            freq[c] = 0
        }
        freq[c] += 1
    }

    for c, val := range freq {
        if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') {
            maxVowelFreq = max(maxVowelFreq, val)
        }else {
            maxConsonantFreq = max(maxConsonantFreq, val)
        }
    }
    return maxVowelFreq + maxConsonantFreq
}