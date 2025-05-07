import java.util.Scanner;

public class Main
{
    static int k;
    static String[] arr;
    static int[] vis;
    static String smin;
    static String smax;
    static Long max = Long.MIN_VALUE;
    static Long min = Long.MAX_VALUE;

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        k = sc.nextInt();
        vis = new int[10];
        arr = new String[k];
        sc.nextLine();
        String input = sc.nextLine();
        arr = input.split(" ");
//        for (int i=0; i<k; i++){
//            arr[i] = sc.nextLine().charAt(0);
//        }

        dfs(0, "", 0);
        System.out.println(smax);
        System.out.println(smin);

    }

    private static void dfs(int depth, String s, int cur) {
        if (depth == k+1){
            try {
                Long x = Long.parseLong(s);
                if (x<min){
                    min = x;
                    smin = s;
                }
                if (x>max){
                    max = x;
                    smax = s;
                }
                return;
            }
            catch (NumberFormatException e){
                System.out.println(s);

            }

        }
        for (int i=0; i<=9; i++){
            if (vis[i]==1)continue;
            if (depth==0){
                vis[i] = 1;
                dfs(depth+1, s+String.valueOf(i), i);
                vis[i] = 0;
            }
            else {
                if(arr[depth-1].equals("<")){
                    if(cur<i){
                        vis[i] = 1;
                        dfs(depth+1, s+String.valueOf(i),i);
                        vis[i] = 0;
                    }
                }
                else{
                    if(cur>i){
                        vis[i] = 1;
                        dfs(depth+1, s+String.valueOf(i),i);
                        vis[i] = 0;
                    }
                }
            }
        }
    }


}
