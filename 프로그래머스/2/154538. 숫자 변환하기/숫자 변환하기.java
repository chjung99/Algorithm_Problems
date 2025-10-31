import java.util.*;

class Solution {
    public int solution(int x, int y, int n) {
        int MAX = 1_000_000;
        Deque<int[]> q = new ArrayDeque<>();
        Set<Integer> visited = new HashSet<>();
        
        q.add(new int[]{x, 0});
        visited.add(x);
        
        while (!q.isEmpty()) {
            int[] cur = q.pollFirst();
            int value = cur[0];
            int cnt = cur[1];
            
            if (value == y){
                return cnt;
            }
            
            int[] nexts = {value + n, value * 2, value * 3};
            for (int next: nexts) {
                if (next > MAX) continue;
                if (!visited.contains(next)){
                    q.add(new int[]{next, cnt + 1});
                    visited.add(next);
                }
            }
        }
        
        return -1;
    }

}