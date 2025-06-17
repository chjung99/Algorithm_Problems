import java.util.Scanner;

public class Main
{
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int[] arr = new int[31];
        for (int i=0; i<28; i++){
            int k = sc.nextInt();
            arr[k] = 1;
        }
        int cnt = 0;
        for (int i=1; i<=30; i++){
            if (cnt == 2){
                break;
            }
            if (arr[i]==0){
                System.out.println(i);
                cnt += 1;
            }
        }
    }

}
