class Solution {
    public List<String> summaryRanges(int[] nums) {
        Deque<Integer> deque = new ArrayDeque<>();
        int n = nums.length;
        List<String> ans = new ArrayList<>();

        for (int i = 0; i < n; i++) {
            if (deque.isEmpty() || deque.peekLast() + 1 == nums[i]) {
                deque.add(nums[i]);
                continue;
            }
            ans.add(getRanges(deque));
            deque.clear();
            deque.add(nums[i]);
        }
        if (deque.size() != 0) {
            ans.add(getRanges(deque));
            deque.clear();
        }
        return ans;
    }

    public String getRanges(Deque<Integer> deque) {
        if (deque.size() == 1) {
            return String.valueOf(deque.pop());
        }
        return deque.peekFirst() + "->" + deque.peekLast();
    }
}