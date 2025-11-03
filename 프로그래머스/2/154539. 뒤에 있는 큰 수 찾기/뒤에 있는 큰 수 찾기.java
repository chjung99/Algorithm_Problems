import java.util.*;

class Solution {
    public int[] solution(int[] numbers) {
        int[] answer = new int[numbers.length];
        int N = numbers.length;
        Deque<Integer> stack = new ArrayDeque<>();
        
        for (int i=0; i<N; i++) answer[i] = -1;
        
        for (int i=0; i<N; i++) {
            while (!stack.isEmpty() && numbers[stack.peek()] < numbers[i]) {
                answer[stack.pop()] = numbers[i];
            }
            stack.push(i);
        }
        
        return answer;
    }
}
