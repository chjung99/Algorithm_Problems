func defangIPaddr(address string) string {
    arr := make([]rune, 0)

    for _, ch := range address {
        if (ch == '.') {
            arr = append(arr, '[')
            arr = append(arr, '.')
            arr = append(arr, ']')
        } else {
            arr = append(arr, ch)
        }
    }
    return string(arr)
}