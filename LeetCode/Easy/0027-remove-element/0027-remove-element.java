class Solution {
    public int removeElement(int[] nums, int val) {
        int n = nums.length;
        int k = n;
        int REMOVED_FLAG = 51;
        for (int i = 0; i < n; i++) {
            if (nums[i] == val) {
                k--;
                nums[i] = REMOVED_FLAG;
            }
        }
        Arrays.sort(nums);
        return k;
}   
}