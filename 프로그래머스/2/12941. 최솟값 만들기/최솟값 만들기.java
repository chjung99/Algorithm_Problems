import java.util.*;

class Solution {
    public int solution(int[] A, int[] B) {
        int answer = 0;

        Arrays.sort(A); // 오름차순
        Arrays.sort(B); // 오름차순

        // A를 뒤집어서 내림차순처럼 사용
        for (int i = 0; i < A.length; i++) {
            answer += A[A.length - 1 - i] * B[i];
        }

        return answer;
    }
}