import java.util.Scanner;
import java.util.ArrayDeque;

public class Main
{
    static int n, L, R;
    static int[][] arr;
    static int[][] visit;
    static int sum;
    static int[] dx = {0, 1, 0, -1};
    static int[] dy = {1, 0, -1, 0};
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		
		int answer = 0;

		n = sc.nextInt();
		L = sc.nextInt();
		R = sc.nextInt();
		
		arr = new int[n][n];
		
		for (int i=0; i<n; i++){
		    for(int j=0; j<n; j++){
		        arr[i][j] = sc.nextInt();
		    }
		}
		
		while (true){
		    int flag = 1;
		    
		    visit = new int[n][n];
		    ArrayDeque<int[]> q;
		    
		    for(int i=0; i<n;i++){
		        for(int j=0; j<n; j++){
		            sum = 0;
		            if (visit[i][j]==1){
		                continue;
		            }
		            
		            q = bfs(i, j);
		          //  for (int k=0; k<n;k++){
		          //      for (int p=0; p<n; p++){
		          //          System.out.print(arr[k][p]+" ");
		          //      }
		          //      System.out.println();
		          //  }
		          //  System.out.println("=========");
		            int cnt = q.size();
		            int num = (int)(sum/cnt);
		            if(cnt==1){
		                continue;
		            }
		            else{
		                flag=0;
		            }
		            while (q.size() !=0){
		                int[] p = q.poll();
		                int x = p[0];
		                int y = p[1];
		                arr[x][y] = num;
		            }
		        }
		    }
		    if (flag==1){
		        break;
		    }
		    answer += 1;
		}
		// 연합 찾기
		
		// 인구 이동
		System.out.print(answer);
    }
    public static ArrayDeque<int[]> bfs(int x, int y){
        ArrayDeque<int[]> ret = new ArrayDeque<>();
        ArrayDeque<int[]> q = new ArrayDeque<>();
        visit[x][y] = 1;
        sum = arr[x][y];
        
        q.add(new int[]{x, y});
        ret.add(new int[]{x, y});
        
        while (q.size()!=0){
            int[] p = q.poll();
            int cx = p[0];
            int cy = p[1];
            
            for (int i=0; i<4; i++){
                int nx = cx + dx[i];
                int ny = cy + dy[i];
                if (outOfRange(nx, ny)||visit[nx][ny] ==1){
                    continue;
                }
                if (Math.abs(arr[nx][ny]-arr[cx][cy])>=L && Math.abs(arr[nx][ny]-arr[cx][cy])<=R){
                    visit[nx][ny] = 1;
                    q.add(new int[]{nx, ny});
                    ret.add(new int[]{nx, ny});
                    sum += arr[nx][ny];
                }
            }
        }
        return ret;
    }
    public static boolean outOfRange(int x, int y){
        return (x < 0 || x>=n || y < 0 || y >= n);
    }
}
