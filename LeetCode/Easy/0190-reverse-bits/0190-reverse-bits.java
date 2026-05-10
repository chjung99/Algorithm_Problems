class Solution {
    public int reverseBits(int n) {
        StringBuilder sb = new StringBuilder();

        while (n != 0) {
            sb.append(Integer.toString(n % 2));
            n = n / 2;
        }

        int padSize = 32 - sb.length();

        for (int i = 0; i < padSize ; i++) {
            sb.append('0');
        }

        String rev = sb.toString();
        int ans = 0;
        int fac = 1;

        for (int i = 31; i >= 0; i--) {
            ans += fac * (rev.charAt(i) - '0');
            fac *= 2;
        }

        return ans;
    }
}