class Solution {
    public boolean isHappy(int n) {
        Set<Integer> visit = new HashSet<>();

        while (true) {
            if (n == 1) {
                break;
            }
            if (visit.contains(n)) {
                return false;
            }
            visit.add(n);
            n = replace(n);
        }
        return true;
    }

    public int replace(int x) {
        int ret = 0;
        while (x != 0) {
            ret += (x % 10) * (x % 10);
            x = x / 10;
        }
        return ret;
    }
}