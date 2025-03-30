import java.util.Scanner;

public class Main
{
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        int A = sc.nextInt();
        int B = sc.nextInt();
        int C = sc.nextInt();
        int N = sc.nextInt();

        for (int i=0; i<= 300; i++){
            for (int j=0; j<= 300; j++){
                for (int k=0; k<= 300 ;k++){
                    if (A*i+B*j+C*k==N){
                        System.out.print("1");
                        return;
                    }
                }
            }
        }
        System.out.print("0");
    }
}
