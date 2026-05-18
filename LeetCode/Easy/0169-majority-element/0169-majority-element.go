func majorityElement(nums []int) int {
    prev := nums[0];
    cnt := 0;
    maxCnt := 0;
    majority := nums[0];
    slices.Sort(nums);
    
    for i := 0; i < len(nums); i++ {
        if (nums[i] == prev) {
            cnt += 1;
        } else {
            prev = nums[i];
            cnt = 1;
        }

        if (maxCnt < cnt ) {
            maxCnt = cnt;
            majority = prev;
        }
    }
    if (maxCnt < cnt ) {
        maxCnt = cnt;
        majority = prev;
    }
    return majority;
}