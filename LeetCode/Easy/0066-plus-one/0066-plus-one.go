func plusOne(digits []int) []int {
    var output []int
    n := len(digits)
    carry := 1
    rest := 0

    for i := n - 1; i >= 0; i-- {
        rest = (digits[i]+carry) % 10
        carry = int((digits[i]+carry) / 10)

        output = append(output, rest)
    }
    if (carry > 0) {
        output = append(output, carry)
    }

    for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
        output[i], output[j] = output[j], output[i]
    }

    return output
}