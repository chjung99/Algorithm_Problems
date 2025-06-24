import java.util.Scanner;

public class Main
{
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int A,B,C,D;
        int H = 0;
        int W = 0;
        A = sc.nextInt();
        B = sc.nextInt();
        C = sc.nextInt();
        D = sc.nextInt();

        for (int i=0; i<24; i++){
            if (H + A <= D){
                H += A;
                W += B;
            }
            else {
                H = Math.max(H-C, 0);

            }
        }
        System.out.println(W);
    }


}
