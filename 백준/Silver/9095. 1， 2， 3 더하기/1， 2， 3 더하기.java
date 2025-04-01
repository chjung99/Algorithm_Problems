import java.util.Scanner;

public class Main
{
    static int answer;
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		int T = sc.nextInt();
		int n;
		for (int i=0; i<T; i++){
		    n = sc.nextInt();
		    answer = 0;
		    dfs(0, n);
		    System.out.println(answer);
		}
		
	}
	static void dfs(int sum, int n){
	    if (sum == n){
	        answer += 1;
	        return;
	    }
	    else if(sum > n){
	        return;
	    }
	    dfs(sum+1, n);
	    dfs(sum+2, n);
	    dfs(sum+3, n);
	}
}
