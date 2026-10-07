func truncateSentence(s string, k int) string {
    end := 0
    cnt := 0

    for i, word := range s {
        if word == ' ' {
            cnt += 1
        }

        if cnt == k {
            end = i
            break
        }
    }
    if end == 0 {
        end = len(s)
    }
    fmt.Print(end)
    return s[:end]
}