import java.util.Scanner;

public class Main
{
	public static void main(String[] args) {
		int N, K;
		int answer = Integer.MIN_VALUE;
		Scanner sc = new Scanner(System.in);
		N = sc.nextInt();
		K = sc.nextInt();
		int[] arr = new int[N];
		
		for (int i=0;i<N;i++){
		    arr[i] = sc.nextInt();
		}
		for (int i=0;i<N-K+1;i++){
		    int sum = 0;
		    for (int j=i;j<i+K;j++){
		        sum += arr[j];
		    }
		    answer = Math.max(answer, sum);
		}
		System.out.print(answer);
	}
}
