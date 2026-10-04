func decode(encoded []int, first int) []int {
    decoded := make([]int, 0)
    decoded = append(decoded, first)

    n := len(encoded)

    for i := 1; i <= n; i++ {
        decoded = append(decoded, encoded[i-1] ^ decoded[i-1])
    }
    return decoded
}