func stableMountains(height []int, threshold int) []int {
    s := make([]int, 0)

    for i, _ := range height {
        if i == 0 {
            continue
        }
        if height[i-1] > threshold {
            s = append(s, i)
        }
    }
    return s
}