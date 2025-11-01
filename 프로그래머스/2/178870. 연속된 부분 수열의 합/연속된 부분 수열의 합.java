class Solution {
    public int[] solution(int[] sequence, int k) {
        int start = 0;
        int end = 0;
        int sum = sequence[end];
        
        int MAX = sequence.length-1;
        
        int answerLeft = -1;
        int answerRight = -1;
        
        while (start <= end) {
            
            if (sum == k) {
                if (answerLeft == -1 && answerRight == -1) {
                    answerLeft = start;
                    answerRight = end;
                } else if (answerRight-answerLeft > end - start) {
                    answerLeft = start;
                    answerRight = end;
                }
                end ++;
                if (end > MAX) break;
                sum += sequence[end];
            } else if (sum < k) {
                end ++;
                if (end > MAX) break;
                sum += sequence[end];
            } else {
                sum -= sequence[start];
                start ++;
            }
        }
        
        return new int[]{answerLeft, answerRight};
    }
}