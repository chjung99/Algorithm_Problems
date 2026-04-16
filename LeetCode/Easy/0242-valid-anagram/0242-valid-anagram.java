class Solution {
    public boolean isAnagram(String s, String t) {
        int SIZE = 26;
        int[] cnt = new int[SIZE];
        int n = s.length();
        int m = t.length();

        if (n != m) {
            return false;
        }

        for (int i = 0; i < n; i++) {
            cnt[s.charAt(i)-'a'] += 1;
        }

        for (int i = 0; i < m; i++) {
            if (cnt[t.charAt(i)-'a'] > 0) {
                cnt[t.charAt(i)-'a'] -= 1;
            } else {
                return false;
            }
        }

        return true;
    }
}