import java.util.Scanner;
import java.util.Stack;
import java.util.Arrays;

public class Main {
    static int[] arr;
    static Stack<Integer> st;
    static int N, M;
    static StringBuilder sb = new StringBuilder(); // 출력용 StringBuilder

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        N = sc.nextInt();
        M = sc.nextInt();
        arr = new int[N];
        st = new Stack<Integer>();

        for (int i = 0; i < N; i++) {
            arr[i] = sc.nextInt();
        }

        Arrays.sort(arr); // 사전순 출력을 위한 정렬
        dfs(0);

        // 최종 출력
        System.out.print(sb.toString());
    }

    public static void dfs(int depth) {
        if (depth >= M) {
            if (st.size() == M) {
                for (int x : st) {
                    sb.append(x).append(" ");
                }
                sb.append("\n");
            }
            return;
        }

        for (int i = 0; i < N; i++) {
            st.push(arr[i]);
            dfs(depth + 1);
            st.pop();
        }
    }
}
