class Solution {
    public int jump(int[] nums) {
        int n = nums.length;
        int[] dist = new int[n];
        for (int i = 0; i < n; i++) {
            dist[i] = Integer.MAX_VALUE;
        }
        
        Deque<Integer> q = new ArrayDeque<>();
        q.add(0);
        dist[0] = 0;

        while (q.size() > 0) {
            int cur = q.poll();

            for (int i = 1; i <= nums[cur]; i++) {
                int nxt = cur + i;
                if (nxt >= n) {
                    continue;
                }
                if (dist[nxt] > dist[cur] + 1) {
                    dist[nxt] = dist[cur] + 1;
                    q.add(nxt);
                }
            }
        }

        return dist[n-1];
    }
}