import java.util.*;

class Solution {
    public int solution(int[] queue1, int[] queue2) {
        int n = queue1.length;
        int totalLen = n * 2;
        
        // queue1 + queue2 합친 배열
        int[] arr = new int[totalLen];
        for (int i = 0; i < n; i++) arr[i] = queue1[i];
        for (int i = 0; i < n; i++) arr[n + i] = queue2[i];
        
        long sum1 = Arrays.stream(queue1).asLongStream().sum();
        long sum2 = Arrays.stream(queue2).asLongStream().sum();
        long total = sum1 + sum2;
        
        // 전체 합이 홀수면 불가능
        if (total % 2 != 0) return -1;
        long target = total / 2;
        
        int left = 0;
        int right = n - 1;
        long curSum = sum1;
        int answer = 0;
        
        // 전체 배열을 최대 3배 길이만큼 순회 (무한 루프 방지)
        int limit = totalLen * 3;
        
        while (answer <= limit) {
            if (curSum == target) return answer;
            
            if (curSum > target) {  // 큐1에서 빼야 함
                curSum -= arr[left % totalLen];
                left++;
            } else {  // 큐2에서 빼서 큐1로 이동
                right++;
                curSum += arr[right % totalLen];
            }
            answer++;
        }
        
        return -1;  // 불가능한 경우
    }
}