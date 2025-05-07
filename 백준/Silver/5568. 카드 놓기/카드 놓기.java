import java.util.Scanner;
import java.util.HashSet;

public class Main
{
    static int n, k;
    static int[] arr;
    static int[] vis;
    static HashSet<String> set;

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        n = sc.nextInt();
        k = sc.nextInt();
        arr = new int[n];
        vis = new int[n];
        set = new HashSet<>();
        for (int i=0; i<n; i++){
            arr[i] = sc.nextInt();
        }

        dfs(0, "");
        System.out.print(set.size());
//        for (String s: set){
//            System.out.print(s+" ");
//        }
    }

    private static void dfs(int depth, String s) {
        if(depth==k){
            set.add(s);
            return;
        }
        for (int i=0; i<n; i++){
            if (vis[i]==0) {
                vis[i] = 1;
                dfs(depth + 1, s + arr[i]);
                vis[i] = 0;
            }
        }
    }

}
