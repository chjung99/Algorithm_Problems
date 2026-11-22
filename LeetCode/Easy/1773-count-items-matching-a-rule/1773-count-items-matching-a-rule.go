func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    cnt := 0

    for _, item := range items {
        t, c, n := item[0], item[1], item[2]
        if ruleKey == "type" && t == ruleValue {
            cnt ++
        } else if ruleKey == "color" && c == ruleValue {
            cnt ++
        } else if ruleKey == "name" && n == ruleValue {
            cnt ++
        }
    }
    return cnt
}