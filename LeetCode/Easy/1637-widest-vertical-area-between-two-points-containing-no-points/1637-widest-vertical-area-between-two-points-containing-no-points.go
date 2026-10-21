func maxWidthOfVerticalArea(points [][]int) int {
    n := len(points)
    maxWidth := 0

    sort.Slice(points[:], func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

    for i := 0; i < n - 1; i++ {
        maxWidth = max(maxWidth, points[i+1][0] - points[i][0])
    }

    return maxWidth
}