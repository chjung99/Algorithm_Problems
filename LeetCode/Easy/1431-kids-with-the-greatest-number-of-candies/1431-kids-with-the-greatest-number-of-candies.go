func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxNum := candies[0]
    ret := make([]bool, 0)

    for _, num := range candies {
        maxNum = max(maxNum, num)
    }

    for _, num := range candies {
        if (num + extraCandies >= maxNum) {
            ret = append(ret, true)
        } else {
            ret = append(ret, false)
        }
    }
    return ret
    
}