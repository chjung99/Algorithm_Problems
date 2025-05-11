import java.util.Scanner;

public class Main
{
    static int N, M;
    static int[][] arr;
    static int ans;
    static int[] dx = {0, 1, 0, -1};
    static int[] dy = {1, 0, -1, 0};
    static int[][] vis;
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        N = sc.nextInt();
        M = sc.nextInt();

        arr = new int[N][M];
        vis = new int[N][M];

        for (int i=0; i<N; i++){
            for (int j=0; j<M; j++){
                arr[i][j] = sc.nextInt();
            }
        }
        for (int i=0; i<N; i++) {
            for (int j = 0; j < M; j++) {
                vis[i][j] = 1;
                dfs(i, j, 0, arr[i][j]);
                vis[i][j] = 0;
            }
        }
        System.out.print(ans);
    }
    public static void dfs(int cx,int cy, int depth, int sum){
        if (depth == 3){
            ans = Math.max(ans, sum);
            return;
        }
        for (int i=0; i<5; i++){
            if (i < 4){

                int nx = cx + dx[i];
                int ny = cy + dy[i];

                if (outOfRange(nx, ny)){
                    continue;
                }
                if (vis[nx][ny] == 1){
                    continue;
                }
                vis[nx][ny] = 1;
                dfs(nx, ny, depth+1, sum+arr[nx][ny]);
                vis[nx][ny] = 0;
            }
            else{
                if (depth == 1){
                    for (int j=0; j<4; j++){
                        int nx = cx + dx[j];
                        int ny = cy + dy[j];
                        if (outOfRange(nx, ny))continue;
                        if (vis[nx][ny]==1) continue;
                        for (int k=0; k<4;k++){
                            if (j==k)continue;
                            int nnx = cx + dx[k];
                            int nny = cy + dy[k];
                            if (outOfRange(nnx, nny))continue;
                            if(vis[nnx][nny]==1)continue;
                            vis[nx][ny] = 1;
                            vis[nnx][nny] = 1;
                            dfs(nnx, nny, depth+2,sum+arr[nx][ny]+arr[nnx][nny]);
                            vis[nx][ny] = 0;
                            vis[nnx][nny] = 0;
                        }
                    }
                }
            }
        }
    }

    private static boolean outOfRange(int nx, int ny) {
        return nx < 0 || nx >= N || ny < 0 || ny >= M;
    }
}
