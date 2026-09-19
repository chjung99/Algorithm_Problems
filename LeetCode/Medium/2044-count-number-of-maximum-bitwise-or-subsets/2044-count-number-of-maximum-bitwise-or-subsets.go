func countMaxOrSubsets(nums []int) int {
    var countNumberOfSubset func(curIdx int, curBitwiseOR int)
    maxBitwiseOR := 0
    numOfSubset := 0
    n := len(nums)

    for _, num := range nums {
        maxBitwiseOR |= num
    }

    countNumberOfSubset = func(curIdx int, curBitwiseOR int) {

        if (curIdx == n) {
            if (curBitwiseOR == maxBitwiseOR) {
                numOfSubset += 1
            }
            return
        }

        countNumberOfSubset(curIdx+1, curBitwiseOR)
        countNumberOfSubset(curIdx+1, curBitwiseOR|nums[curIdx])

    }

    countNumberOfSubset(0, 0)

    return numOfSubset
}