class Solution {
    public void rotate(int[] nums, int k) {
        int n = nums.length;
        int[] rottated = new int[n];

        for (int i = 0; i < n; i++) {
            rottated[(i+k)%n] = nums[i];
        }

        for (int i = 0; i < n; i++) {
            nums[i] = rottated[i];
        }
    }
}