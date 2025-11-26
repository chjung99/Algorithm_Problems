class Solution {
    int SIZE = 26;
    public boolean isAnagram(String s, String t) {
        int[] countS = new int[SIZE];

        if (s.length() != t.length()) return false;
        
        for (int i = 0; i < s.length(); i++){
            countS[s.charAt(i)-'a'] += 1;
            countS[t.charAt(i)-'a'] -= 1;
        }
        for (int i = 0; i < SIZE; i++){
            if (countS[i] != 0) return false;
        }
        return true;
    }
}