func addBinary(a string, b string) string {
    var r []byte
    padA, padB := addPadding(a, b)
    revA, revB := reverse(padA), reverse(padB)
    
    n := len(padA)
    var carry byte
    var rest byte

    for i := 0; i < n; i++ {
        rest = getRest(carry, revA[i], revB[i])
        carry = getCarry(carry, revA[i], revB[i])
        r = append(r, rest)
    }
    if (carry == '1') {
        r = append(r, '1')
    }

    return reverse(string(r))
}

func getSetCount(carry byte, x byte, y byte) int {
    cnt := 0
    if (carry == '1') {
        cnt += 1
    }
    if (x == '1') {
        cnt += 1
    }
    if (y == '1') {
        cnt += 1
    }
    return cnt
}

func getRest(carry byte, x byte, y byte) byte {
    cnt := getSetCount(carry, x, y)
    if (cnt % 2 == 0) {
        return '0'
    } else {
        return '1'
    }
}

func getCarry(carry byte, x byte, y byte) byte {
    cnt := getSetCount(carry, x, y)
    if (int(cnt/2) == 0) {
        return '0'
    } else {
        return '1'
    }
}

func reverse(a string) (string) {
    arrA := []byte(a)

    for i, j := 0, len(a)-1; i < j; i, j = i + 1, j - 1 {
        arrA[i], arrA[j] = arrA[j], arrA[i]
    }
    return string(arrA)
}

func addPadding(a string, b string) (string, string) {
    diff := len(a) - len(b)
    if (diff == 0) {
        return a, b
    } else if (diff > 0) {
        var r []byte
        for i := 0; i < len(a); i++ {
            if (i < diff) {
                r = append(r, '0')
            } else {
                r= append(r, b[i-diff])
            }
        }
        return a, string(r)
    } else {
        diff *= -1
        var r []byte
        for i := 0; i < len(b); i++ {
            if (i < diff) {
                r = append(r, '0')
            } else {
                r = append(r, a[i-diff])
            }
        }
        return string(r), b
    }
    
}