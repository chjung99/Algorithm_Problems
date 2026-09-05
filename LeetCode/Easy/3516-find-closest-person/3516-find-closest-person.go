func findClosest(x int, y int, z int) int {
    distX := max(z-x, x-z)
    distY := max(z-y, y-z)

    if (distX == distY) {
        return 0
    } else if (distX < distY) {
        return 1
    } else {
        return 2
    }
}