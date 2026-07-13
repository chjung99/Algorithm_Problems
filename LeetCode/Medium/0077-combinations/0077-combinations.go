func combine(n int, k int) [][]int {
    ret := make([][]int, 0)

    getCombinationsOf(n, k, 0, 0, []int{}, &ret)

    return ret
}

func getCombinationsOf(n int, k int, prev int, depth int, tmp []int, ret *[][]int) {
    if (depth == k) {
        copied := make([]int, len(tmp))
        copy(copied, tmp)
        (*ret) = append(*ret, copied)
        return
    }

    for i := prev + 1; i <= n; i++ {
        tmp = append(tmp, i)
        getCombinationsOf(n, k, i, depth+1, tmp, ret)
        tmp = tmp[:len(tmp)-1]
    }
}

