class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> location  = new HashMap<>();
        int n = nums.length;
        int[] ans = new int[2];

        for (int i = 0; i < n; i++) {
            location.put(nums[i], i);
        }

        for (int i = 0; i < n; i++) {
            int key = target - nums[i];
            if (location.containsKey(key) && location.get(key) != i) {
                ans[0] = i;
                ans[1] = location.get(key);
                break;
            }
        }
        return ans;
    }
}