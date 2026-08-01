func scoreOfString(s string) int {
    score := 0
    for i := 0; i < len(s) - 1; i ++ {
        score += max(int(s[i+1])-int(s[i]), int(s[i])-int(s[i+1]))
    }
    return score
}