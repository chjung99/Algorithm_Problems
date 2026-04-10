class Solution {
    public String longestCommonPrefix(String[] strs) {
        String prefix = "";
        int n = strs.length;
        boolean flag = false;
        int k = strs[0].length();
        int idx = 0;

        for (int i = 0; i < n; i++) {
            k = Math.min(k, strs[i].length());
        }
        

        while (true) {
            if (idx == k) {
                break;
            }
            Character c = strs[0].charAt(idx);

            for (int i = 1; i < n; i++) {
                if (c != strs[i].charAt(idx)){
                    flag = true;
                    break;
                }
            }
            if (flag) {
                break;
            }
            idx++;
        }
        return strs[0].substring(0, idx);
        
    }
}