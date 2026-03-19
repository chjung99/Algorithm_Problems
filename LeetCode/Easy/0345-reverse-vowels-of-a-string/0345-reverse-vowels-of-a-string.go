func reverseVowels(s string) string {
var v []rune
	var ret string

	for _, r := range s {
		if (r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U') {
			v = append(v, r)
		}
	}
	n := len(v)
	for i := 0; i < int(n/2); i++ {
		v[i], v[n-i-1] = v[n-i-1], v[i] 
	}
	cnt := 0
	for _, r := range s {
		if (r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U') {
			ret += string(v[cnt])
			cnt += 1
		} else {
			ret += string(r)
		}
	}

	return ret
}