func reversePrefix(word string, ch byte) string {
    start := 0
    end := 0

    arr := make([]rune, 0)

    for _, char := range word {
        arr = append(arr, char)
    }

    for i, char := range word {
        if byte(char) == ch {
            end = i
            break
        }
    }

    for i, j := start, end; i < j; i, j = i + 1, j - 1 {
        arr[i], arr[j] = arr[j], arr[i]
    }

    return string(arr)
}