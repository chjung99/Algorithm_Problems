func simplifyPath(path string) string {
    st := make([]rune, 0)
    dotCnt := 0
    flag := false

    for _, r := range(path) {

        if (len(st) == 0) {
            st = append(st, r)
        } else if (unicode.IsLetter(r) || unicode.IsDigit(r)||r == '_') {
            st = append(st, r)
            flag = true
        } else if (r == '/') {
            if (!flag && dotCnt == 2) {
                for (st[len(st)-1] != '/') {
                    st = st[:len(st)-1]
                }
                if (len(st) != 1) {
                    st = st[:len(st)-1]
                    for (st[len(st)-1] != '/') {
                        st = st[:len(st)-1]
                    }
                }
            } else if (!flag && dotCnt == 1) {
                for (st[len(st)-1] != '/') {
                    st = st[:len(st)-1]
                }
            } else if (st[len(st)-1] == '/') {
                continue
            } else {
                st = append(st, r)
            }
            dotCnt = 0
            flag = false
        } else if (r == '.') {
            st = append(st, r)
            dotCnt += 1
        }
    }

    if (!flag && dotCnt == 2) {
        for (st[len(st)-1] != '/') {
            st = st[:len(st)-1]
        }
        if (len(st) != 1) {
            st = st[:len(st)-1]
            for (st[len(st)-1] != '/') {
                st = st[:len(st)-1]
            }
        }
    } else if (!flag && dotCnt == 1) {
        for (st[len(st)-1] != '/') {
            st = st[:len(st)-1]
        }
    }

    if (len(st) > 1 && st[len(st)-1] == '/') {
        st = st[:len(st)-1]
    }


    return string(st)
}