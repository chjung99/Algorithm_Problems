class Solution {
    public int canCompleteCircuit(int[] gas, int[] cost) {
        int n = gas.length;
        int s = 0;

        for (int i = 0; i < n; i++) {
            s += gas[i] - cost[i];
        }
        
        if (s < 0) {
            return -1;
        }
        
        int idx = 0;
        int curGas = 0;

        for (int i = 0; i < n; i++) {
            if (i > 0) {
                curGas -= cost[i-1];
            }
            if (curGas < 0) {
                idx = i;
                curGas = 0;
            }
            curGas += gas[i];
        }
        return idx;
    }
}