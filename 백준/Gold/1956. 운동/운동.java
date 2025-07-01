/******************************************************************************

 Welcome to GDB Online.
 GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
 C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
 Code, Compile, Run and Debug online from anywhere in world.

 *******************************************************************************/
import java.util.Scanner;
public class Main
{
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int V = sc.nextInt();
        int E = sc.nextInt();
        int[][] dist = new int[V+1][V+1];

        for (int i=0;i<V+1;i++){
            for(int j=0;j<V+1;j++){
                dist[i][j] = Integer.MAX_VALUE;
            }
        }


        for (int i=0;i<E;i++){
            int a = sc.nextInt();
            int b = sc.nextInt();
            int c = sc.nextInt();
            dist[a][b] = c;
        }

        for (int k=1;k<V+1;k++){
            for(int i=1;i<V+1;i++){
                for(int j=1;j<V+1;j++){
                    if (dist[i][k] != Integer.MAX_VALUE && dist[k][j] != Integer.MAX_VALUE) {
                        dist[i][j] = Math.min(dist[i][j], dist[i][k] + dist[k][j]);
                    }
                }
            }
        }

        int answer = Integer.MAX_VALUE;
        for(int i=0;i<V+1;i++){
            answer = Math.min(dist[i][i], answer);

        }
        if(answer == Integer.MAX_VALUE){
            System.out.print("-1");
        }
        else{
            System.out.print(answer);
        }

    }
}
