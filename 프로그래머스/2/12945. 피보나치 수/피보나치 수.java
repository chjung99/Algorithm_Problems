class Solution {
    public int solution(int n) {
        int R = 1234567;
        int[] F = new int[100_001];
        F[0] = 0;
        F[1] = 1;
        
        for (int i=2; i<=n; i++) {
            F[i] = (F[i-1] + F[i-2])%R;
        }
        
        return F[n];
    }
}