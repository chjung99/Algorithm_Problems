func evalRPN(tokens []string) int {
    st := make([]int, 0)

    for i := 0; i < len(tokens); i++ {
        if (tokens[i] == "+" || tokens[i] == "-" ||tokens[i] == "*" || tokens[i] == "/") {
            operate(tokens[i], &st)
        } else {
            val, err := strconv.Atoi(tokens[i])
            if (err == nil) {
                st = append(st, val)
            }
        }
    }
    return st[0]
}

func operate(operator string, st *[]int) {
    top := (*st)[len(*st)-1]
    (*st) = (*st)[:len(*st)-1]
    if (operator == "+") {
        (*st)[len(*st)-1] += top
    } else if (operator == "-") {
        (*st)[len(*st)-1] -= top
    } else if (operator == "*") {
        (*st)[len(*st)-1] *= top
    } else {
        (*st)[len(*st)-1] /= top
    }
}