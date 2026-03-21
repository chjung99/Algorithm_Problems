func findComplement(num int) int {
    // 0101
    // 0111
    // 0010
    bitLen := bits.Len(uint(num))
    mask := (1 << bitLen) - 1
    return num ^ mask
}