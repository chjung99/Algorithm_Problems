import java.util.Scanner;
import java.util.ArrayDeque;
public class Main
{
	public static void main(String[] args) {
	    Scanner sc = new Scanner(System.in);
        int N = sc.nextInt();
        int M = sc.nextInt();
        int[] A = new int[N];
        int sum = 0;
        int answer = 0;
        
        
        for(int i = 0; i < N; i++){
            A[i] = sc.nextInt();
        }
        
        int head = 0;
        ArrayDeque<Integer> deque = new ArrayDeque<>();
        
        deque.addFirst(A[head]);
        sum = A[head];
        
        while (head < N){
            
            if (sum == M){
                answer ++;
                head += 1;
                if (head < N){
                    deque.addFirst(A[head]);
                    sum += A[head];
                }
            }
            else if(sum > M){
                sum -= deque.pollLast();
            }
            else{
                head += 1;
                if (head < N){
                    deque.addFirst(A[head]);
                    sum += A[head];
                }
            }
        }
        System.out.print(answer);
        
	}
	
}
