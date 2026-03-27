class Solution {
    public int removeDuplicates(int[] nums) {
        int n = nums.length;
        int k = n;
        int FLAG = 100000;
        for (int i = n-1; i >= 2; i--) {
            if (nums[i] == nums[i-2]) {
                nums[i] = FLAG;
                k --;
            }
        }
        int x = 0;
        int t = 0;

        while (true) {
            if (x == k) {
                break;
            }
            if (nums[t] == FLAG) {
                t++;
                continue;
            }
            nums[x] = nums[t];
            x++;
            t++;
        }
        return k;
    }
}