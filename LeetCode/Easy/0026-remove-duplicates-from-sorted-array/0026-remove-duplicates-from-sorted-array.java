class Solution {
    public int removeDuplicates(int[] nums) {
        HashSet visit = new HashSet<>();
        int n = nums.length;
        
        int k = 0;
        int x = 0;

        while (x < n) {
            if (!visit.contains(nums[x])) {
                visit.add(nums[x]);
                nums[k] = nums[x];
                k++;
            }
            x++;
        }

        return k;
    }

}