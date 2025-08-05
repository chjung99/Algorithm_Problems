import java.util.ArrayDeque;
import java.util.Collection;
import java.util.Queue;
import java.util.Scanner;

public class Main {
    static int[] visit;
    static int[] ans;
    static int[][] adj;
    static int N;
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        N = sc.nextInt();
        adj =  new int[N][N];

        for (int i=0; i<N; i++){
            for (int j=0; j<N; j++){
                adj[i][j]= sc.nextInt();
            }
        }
        for (int i=0; i<N; i++){
            visit = new int[N];
            ans = new int[N];

            bfs(i);
            for (int j=0; j<N; j++){
                System.out.print(ans[j]+" ");
            }
            System.out.println();
        }
    }

    private static void bfs(int start) {
        visit[start] = 1;
        Queue<Integer> queue = new ArrayDeque<>();
        queue.add(start);

        while (!queue.isEmpty()){
            int cur = queue.poll();
            for (int i=0; i<N; i++){
                //길이 없다면 패스
                if(adj[cur][i]==0){
                    continue;
                }
                //이미 방문한적있고, 시작점이 아니라면 패스
                else if (visit[i]==1 && i!=start) {
                    continue;
                }
                //이미 방문한적있고, 시작점이라면
                else if (visit[i]==1 && i==start) {
                    ans[i]=1;
                }
                //방문한 적이 없다면
                else {
                    visit[i] = 1;
                    ans[i] = 1;
                    queue.add(i);
                }
            }
        }
    }
}