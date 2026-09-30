func interpret(command string) string {
    ret := ""
    st := ""
    for _, c := range command {
        if c == 'G' {
            ret += "G"
        } else {
            if c == '(' {
                st += string(c)
            } else if c == ')' {
                tmp := ""
                for len(st) != 0 {
                    if (st[len(st)-1] != '(') {
                        tmp = string(st[len(st)-1]) + tmp
                    }
                    st = st[:len(st)-1]
                }
                if (len(tmp) == 0) {
                    ret += "o"    
                } else {
                    ret += tmp
                }
            } else {
                st += string(c)
            }
        }
    }
    return ret
}