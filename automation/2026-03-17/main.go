func convertToTitle(columnNumber int) string {
	var rev []byte
	var ans []byte
	m := map[int]byte{}
	SIZE := 26

	for i := 0; i < SIZE; i++ {
		m[i] = byte('A' + i)
	}

	for columnNumber > 0 {
		r := (columnNumber - 1) % SIZE
		rev = append(rev, m[r])
		columnNumber = (columnNumber - 1) / SIZE
	}

	for len(rev) > 0 {
		ans = append(ans, rev[len(rev)-1])
		rev = rev[:len(rev)-1]
	}

	return string(ans)
}