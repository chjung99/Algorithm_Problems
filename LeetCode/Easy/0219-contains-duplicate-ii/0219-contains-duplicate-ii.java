class Solution {
    public boolean containsNearbyDuplicate(int[] nums, int k) {
        Map<Integer, Integer> map = new HashMap<>();
        int n = nums.length;

        for (int i = 0; i < n; i++) {
            int loc = map.getOrDefault(nums[i], -1);
            if (loc == -1) {
                map.put(nums[i], i);
                continue;
            }

            if (i - loc <= k) {
                return true;
            }
            map.put(nums[i], i);
        }
        return false;
    }
}