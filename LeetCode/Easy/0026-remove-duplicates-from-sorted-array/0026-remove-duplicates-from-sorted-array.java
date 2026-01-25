class Solution {
    public int removeDuplicates(int[] nums) {
        Set<Integer> set = new HashSet<>();
        for (int num: nums) {
            set.add(num);
        }
        int k = set.size();
        List<Integer> ans = new ArrayList<>(set);
        Collections.sort(ans);
        for (int i = 0; i < k; i++) {
            nums[i] = ans.get(i);
        }
        return k;
    }
}