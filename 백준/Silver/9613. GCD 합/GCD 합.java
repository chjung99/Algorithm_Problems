import java.util.Scanner;
import java.util.Arrays;
public class Main
{
	public static void main(String[] args) {
		int t, n;
		long sum;
		int[] arr;
		Scanner sc = new Scanner(System.in);
		t = sc.nextInt();
		for (int i=0; i<t; i++){
		    n = sc.nextInt();
		    sum = 0;
		    arr = new int[n];
		    for (int j=0; j<n; j++){
		        arr[j] = sc.nextInt();
		    }
		  //  Arrays.sort(arr);
		    for (int p=0; p<n; p++){
		        for (int q=p+1; q<n; q++){
		            if (arr[p]>arr[q]){
		                sum += calcGCD(arr[p], arr[q]);
		            }
		            else{
		                sum += calcGCD(arr[q], arr[p]);
		            }
		        }
		    }
		    System.out.println(sum);
		    
		}
		
	}
	public static int calcGCD(int x, int y){
	    if (y==0) return x;
	    return calcGCD(y, x % y);
	}
}
