class Solution {
    public int[] countBits(int n) {
        int[] ans = new int[n+1];

        for (int i = 0; i <= n; i++) {
            ans[i] = numberOfSetBit(i);
        }

        return ans;
    }

    public int numberOfSetBit(int n) {
        int cnt = 0;
        while (n != 0) {
            if (n % 2 == 1) cnt ++;
            n = (int) n / 2;
        }
        return cnt;
    }
}