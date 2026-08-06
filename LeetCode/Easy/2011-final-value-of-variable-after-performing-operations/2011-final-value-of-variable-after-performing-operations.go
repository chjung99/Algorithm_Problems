func finalValueAfterOperations(operations []string) int {
    val := 0
    for _, s := range operations {
        if (s[0] == '+' || s[len(s)-1] == '+') {
            val += 1
        } else if (s[0] == '-' || s[len(s)-1] == '-'){
            val -= 1
        }
    }
    return val
}