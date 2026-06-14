func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)

    sumGas := 0
    sumCost := 0
    
    idx := 0
    curGas := 0

    for _, num := range gas {
        sumGas += num
    }

    for _, num := range cost {
        sumCost += num
    }

    if (sumGas < sumCost) {
        return -1
    }

    for i := 0; i < n; i++ {
        if (curGas < 0) {
            curGas = 0
            idx = i
        }
        curGas += gas[i]
        curGas -= cost[i]
    }
    return idx
}