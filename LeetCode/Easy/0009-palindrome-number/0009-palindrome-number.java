class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0) {
            return false;
        }

        int src = x;
        int rev = 0;

        while (x != 0) {
            rev *= 10;
            rev += x % 10;
            x = x/10;
        }
        return src == rev;
    }
}