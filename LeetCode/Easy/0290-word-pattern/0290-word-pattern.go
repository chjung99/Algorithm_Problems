func wordPattern(pattern string, s string) bool {
	n := len(pattern)
	words := strings.Fields(s)
    if (n != len(words)) {
        return false
    }
	m := make(map[byte]string)
	r := make(map[string]byte)
	
	for i := 0; i < n; i++ {
		cur := pattern[i]

		_, mexists := m[cur]
		_, rexists := r[words[i]]
		if (!mexists && !rexists) {
			m[cur] = words[i]
			r[words[i]] = cur
		} else if (rexists && r[words[i]] != cur) {
			return false
		} else if (m[cur] != words[i]){
			return false
		}
	}
	return true
}