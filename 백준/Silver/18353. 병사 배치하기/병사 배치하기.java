import java.util.Scanner;

public class Main
{
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		
		int n = sc.nextInt();
		int[] dp = new int[n];
		int[] arr = new int[n];
		for (int i=0;i<n;i++){
		    arr[i] = sc.nextInt();
		}
		for (int i=0;i<n;i++){
		    dp[i] = 1;
		}
		int maxValue = Integer.MIN_VALUE;
		for(int i=0; i<n ; i++){
		    for (int j=0;j<i;j++){
		        if(arr[j] > arr[i]){
		            dp[i] = Math.max(dp[i], dp[j]+1);
		            
		        }
		    }
		}
		for(int i=0;i<n;i++){
		    maxValue = Math.max(maxValue, dp[i]);
		}
		System.out.print(n-maxValue);
	}
	
}
