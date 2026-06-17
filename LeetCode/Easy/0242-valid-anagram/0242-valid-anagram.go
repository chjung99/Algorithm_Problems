func isAnagram(s string, t string) bool {
    smap := make(map[rune]int)
    tmap := make(map[rune]int)

    if (len(s) != len(t)) {
        return false
    }

    for _, c := range s {
        _, exist := smap[c]
        if (!exist) {
            smap[c] = 1
        } else{
            smap[c] += 1
        }
    }

    for _, c := range t {
        _, exist := tmap[c]
        if (!exist) {
            tmap[c] = 1
        } else{
            tmap[c] += 1
        }
    }

    for key, value := range smap {
        _, exists := tmap[key]
        if (!exists) {
            return false
        }
        if (value != tmap[key]) {
            return false
        }
    }

    return true
}