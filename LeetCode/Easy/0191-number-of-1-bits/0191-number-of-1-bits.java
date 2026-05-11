class Solution {
    public int hammingWeight(int n) {
        int cnt = 0;
        while (n != 0) {
            cnt += n % 2;
            n = n / 2;
        }
        return cnt;
    }
}