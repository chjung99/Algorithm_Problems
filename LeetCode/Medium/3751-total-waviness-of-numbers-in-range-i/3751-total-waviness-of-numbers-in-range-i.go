func totalWaviness(num1 int, num2 int) int {
    cnt := 0

    for i := num1; i <= num2; i++ {
        cnt += countWaviness(i)
    }
    return cnt
}

func countWaviness(num int) int {
    cnt := 0
    arr := make([]int, 0)

    for num != 0 {
        arr = append(arr, num % 10)
        num = num / 10
    }

    slices.Reverse(arr)

    for i := 1; i < len(arr) - 1; i++ {
        if (arr[i] > arr[i-1] && arr[i] > arr[i+1]) {
            cnt += 1
        }
        if (arr[i] < arr[i-1] && arr[i] < arr[i+1]) {
            cnt += 1
        }
    }
    return cnt
}