import java.util.*;

class Solution {
    public int solution(int[] A, int[] B) {
        int answer = 0;

        // 1️⃣ A 오름차순 정렬
        Arrays.sort(A);
        // 2️⃣ 배열 뒤집어서 내림차순으로 사용
        for (int i = 0; i < A.length / 2; i++) {
            int temp = A[i];
            A[i] = A[A.length - 1 - i];
            A[A.length - 1 - i] = temp;
        }

        // 3️⃣ B 오름차순 정렬
        Arrays.sort(B);

        // 4️⃣ 곱셈 합산
        for (int i = 0; i < A.length; i++) {
            answer += A[i] * B[i];
        }

        return answer;
    }
}