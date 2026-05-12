class Solution {
    public static int minSubArrayLen(int target, int[] nums) {
        Deque<Integer> q = new ArrayDeque<>();
        int n = nums.length;
        int len = n + 1;
        int sum = nums[0];

        int cur = 0;
        q.add(nums[0]);

        while (cur < n) {
            if (sum == target) {
                len = Math.min(len, q.size());
                sum -= q.getFirst();
                q.removeFirst();
                cur += 1;
                if (cur < n) {
                    q.addLast(nums[cur]);
                    sum += nums[cur];
                }
            } else if (sum > target) {
                len = Math.min(len, q.size());
                sum -= q.getFirst();
                q.removeFirst();
            } else {
                cur += 1;
                if (cur < n) {
                    q.addLast(nums[cur]);
                    sum += nums[cur];
                }
            }

        }
        if (len == n + 1) {
            return 0;
        }
        return len;
    }
}