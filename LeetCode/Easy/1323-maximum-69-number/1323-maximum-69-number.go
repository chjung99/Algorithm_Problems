func maximum69Number (num int) int {
    digits := make([]int, 0)
    
    for num != 0 {
        digits = append(digits, num % 10)
        num = num / 10
    }
    slices.Reverse(digits) 

    for i, digit := range digits {
        if digit == 6 {
            digits[i] = 9
            break
        }
    }
    for _, digit := range digits {
        num *= 10
        num += digit
    }
    return num
}