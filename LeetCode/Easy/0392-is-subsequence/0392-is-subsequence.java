class Solution {
    public boolean isSubsequence(String s, String t) {
        int sLen = s.length();
        int tLen = t.length();

        int sIdx = 0;
        int tIdx = 0;

        while (true) {
            if (sIdx >= sLen || tIdx >= tLen) {
                break;
            }

            if (s.charAt(sIdx) == t.charAt(tIdx)) {
                sIdx ++;
            }

            tIdx++;
        }
        return (sIdx == sLen);
    }
}