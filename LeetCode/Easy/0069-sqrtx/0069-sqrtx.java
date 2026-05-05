class Solution {
    public int mySqrt(int x) {
        if (x < 2) return x;

        int left = 1;
        int right = x / 2; // The square root of x (where x > 1) is never greater than x/2
        int ans = 0;

        while (left <= right) {
            // Standard way to find mid without overflow
            int mid = left + (right - left) / 2;
            
            // Use long to prevent overflow during multiplication
            long num = (long)mid * mid;

            if (num == x) {
                return mid;
            } else if (num < x) {
                ans = mid; // Potential answer, keep looking for a larger one
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        
        return ans;
    }
    
}