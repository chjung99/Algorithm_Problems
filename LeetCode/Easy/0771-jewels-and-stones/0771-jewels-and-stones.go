func numJewelsInStones(jewels string, stones string) int {
    cnt := 0
    for _, stone := range stones {
        for _, jewel := range jewels {
            if (stone == jewel) {
                cnt += 1
            }
        }
    }
    return cnt
}