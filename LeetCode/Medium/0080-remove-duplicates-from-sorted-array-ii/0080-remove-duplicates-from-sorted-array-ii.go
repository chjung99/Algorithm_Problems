func removeDuplicates(nums []int) int {
    prev := nums[0]
    read := 0
    write := 0
    cnt := 0

    n := len(nums)

    for (read < n) {

        if (prev == nums[read]) {
            if (cnt <= 1) {
                nums[write] = nums[read]
                write += 1
                read += 1
            } else {
                read += 1
            }
            cnt += 1

        } else {
            prev = nums[read]
            nums[write] = nums[read]
            write += 1
            read += 1
            cnt = 1
        }
    }

    return write
}