import java.util.Scanner;
import java.util.HashSet;

public class Main
{
    static int answer;
    static int N;
    static int[][] S;
    static HashSet<Integer> startTeam;
    static HashSet<Integer> linkTeam;
    
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        
        answer = Integer.MAX_VALUE;
        
        startTeam = new HashSet<>();
        linkTeam = new HashSet<>();

        N = sc.nextInt();
        
        S = new int[N][N];
        
        for (int i=0; i<N; i++){
            for (int j=0; j<N;j++){
                S[i][j] = sc.nextInt();        
            }
        }
        
        dfs(0);
        
        System.out.print(answer);
    }
    public static void dfs(int depth){
        if (depth == N){
            if (startTeam.size() != 0 && linkTeam.size() != 0){
                int startSum = getSum(startTeam);
                int linkSum = getSum(linkTeam);
                int diff = Math.abs(startSum-linkSum);
                answer = Math.min(answer, diff);
            }
            return;
        }
        //add to startTeam
        startTeam.add(depth);
        dfs(depth+1);
        startTeam.remove(depth);
        //add to linkTeam
        linkTeam.add(depth);
        dfs(depth+1);
        linkTeam.remove(depth);
        
    }
    public static int getSum(HashSet<Integer> team){
        int sum = 0;
        for (int i : team){
            for (int j=0; j<N;j++){
                if (i != j){
                    if (team.contains(j)){
                        sum += S[i][j];
                    }
                }
            }
        }
        return sum;
    }
}
