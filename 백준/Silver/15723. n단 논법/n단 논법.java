import java.util.Scanner;
import java.util.ArrayDeque;

public class Main
{
    static int n;
    static int[] adj;
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		
		n = sc.nextInt();
		adj = new int[100];
		
		sc.nextLine();
		for (int i=0; i<100;i++){
		    adj[i] = -1;
		}
		for (int i=0; i < n; i++){
		    String s = sc.nextLine();
		    
		    int x = (int) s.charAt(0)-'a';
		    int y = (int) s.charAt(5)-'a';
		    
		    adj[x] = y;
		}
		
		int m = sc.nextInt();
		sc.nextLine();
		
		for (int i=0; i < m; i++){
		    String s = sc.nextLine();
		    int x = (int) s.charAt(0)-'a';
		    int y = (int) s.charAt(5)-'a';
		    System.out.println(bfs(x, y));
		}
	}
	public static char bfs(int s, int e){
	    
	    int[] visit = new int[100];
	    ArrayDeque<Integer> q = new ArrayDeque<>();
	    visit[s] = 1;
	    q.add(s);
	    
	    while (q.size()!=0){
	        int cur = q.removeFirst();
	        int nxt = adj[cur];
	        
	        if(nxt==-1||visit[nxt]==1){
	            continue;
	        }
	        visit[nxt]=1;
	        q.add(nxt);
	    }
	    
	    if (visit[e]==1){
	        return 'T';
	    }
	    else{
	        return 'F';
	    }
	}
}
