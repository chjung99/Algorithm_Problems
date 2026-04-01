class Solution {
    public boolean canJump(int[] nums) {
        int n = nums.length;
        boolean[] visit = new boolean[n];

        bfs(nums, visit, n);
        return visit[n-1];
    }

    public void bfs(int[] nums, boolean[] visit, int n) {
        Deque<Integer> q = new ArrayDeque<>();
        visit[0] = true;
        q.add(0);

        while (q.size() > 0) {
            int cur = q.poll();
            for (int i = 1; i <= nums[cur]; i++) {
                int nxt = cur + i;
                if (nxt >= n) {
                    continue;
                }
                if (!visit[nxt]) {
                    visit[nxt] = true;
                    q.add(nxt);
                }
            }
        }
    }
}