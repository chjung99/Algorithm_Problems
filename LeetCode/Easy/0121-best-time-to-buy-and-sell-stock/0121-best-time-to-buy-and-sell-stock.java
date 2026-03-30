class Solution {
    public int maxProfit(int[] prices) {
        int n = prices.length;
        int ans = 0;
        int minValue = prices[0];
        
        for (int i = 0; i < n; i++) {
            ans = Math.max(ans, prices[i] - minValue);
            minValue = Math.min(minValue, prices[i]);
        }

        return ans;
    }
}