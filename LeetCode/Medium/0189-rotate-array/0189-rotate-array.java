class Solution {
    public void rotate(int[] nums, int k) {
        int n = nums.length;
        if (k > n) {
            k = k % n;
        }
        reverse(nums, 0, n-1);
        reverse(nums, 0, k-1);
        reverse(nums, k, n-1);
    }
    public void reverse(int[] nums, int s, int e) {
        int l = e-s +1;
        for (int i = 0; i < l/2; i++) {
            int temp = nums[s+i];
            nums[s+i] = nums[s+l-1-i];
            nums[s+l-1-i] = temp;
        }
    }
}