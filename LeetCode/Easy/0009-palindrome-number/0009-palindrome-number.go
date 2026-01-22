func isPalindrome(x int) bool {
    var arr []int
    num := x
    if (num < 0) {
        return false;
    }
    for num !=0 {
        arr = append(arr, num % 10)
        num = num / 10
    }
    fac := 1
    rev := 0
    for i := len(arr) -1; i>=0; i-- {
        rev += fac * arr[i]
        fac *= 10;
    }
    return x == rev
}