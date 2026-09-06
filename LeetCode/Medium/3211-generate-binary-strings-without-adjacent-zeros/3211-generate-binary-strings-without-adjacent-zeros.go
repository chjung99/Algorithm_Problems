func validStrings(n int) []string {
    ret := make([]string, 0)

    var generateString func(depth int, substring string)

    // 2. 익명 함수를 변수에 할당합니다.
    generateString = func(depth int, substring string) {
        if depth == n {
            // string은 불변 객체이므로 copy 없이 그냥 append 하면 됩니다.
            ret = append(ret, substring)
            return
        }
        
        if depth == 0 || substring[depth-1] == '1' {
            generateString(depth+1, substring+"0")
            generateString(depth+1, substring+"1")
        } else {
            generateString(depth+1, substring+"1")
        }
    }

    generateString(0, "")

    return ret
}