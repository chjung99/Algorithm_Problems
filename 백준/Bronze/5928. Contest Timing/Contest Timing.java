import java.util.*;
public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int D, H, M;

        D = sc.nextInt();
        H = sc.nextInt();
        M = sc.nextInt();

        int E = (D-11)*60*24+(H-11)*60+(M-11);
        if (E < 0){
            E = -1;
        }
        System.out.println(E);
    }
}