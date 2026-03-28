class Solution {
    public int majorityElement(int[] nums) {
        int n = nums.length;
        int m = (int)(n/2+1);

        Arrays.sort(nums);

        int x = nums[0];
        int cnt = 0;
        // System.out.println(m);
        for (int i = 0; i < n; i++) {
            // System.out.println(x +", " + cnt);
            if (x == nums[i]) {
                cnt ++;
            }
            if (cnt == m) {
                break;
            }
            if (x != nums[i]) {
                x = nums[i];
                cnt = 1;
            }
        }
        
        return x;
    }
}