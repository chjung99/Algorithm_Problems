/******************************************************************************

 Welcome to GDB Online.
 GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
 C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
 Code, Compile, Run and Debug online from anywhere in world.

 *******************************************************************************/
import java.util.Scanner;
import java.util.HashSet;

public class Main
{

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int N = sc.nextInt();
        int[] arr = new int[N];
        HashSet<Integer> triNums = new HashSet<Integer>();
        for(int i=0;i<N;i++){
            arr[i] = sc.nextInt();
        }
        for(int i=1;i<= 1000;i++){
            triNums.add(i*(i+1)/2);
        }
        for(int i=0;i<N;i++){
            int k = arr[i];
            int flag = 0;
            for(int p=0;p<k;p++){
                for(int q=0;q<k;q++){
                    if(p==q){
                        continue;
                    }
                    int[] nums = calculateNum(k, p, q);
                    if (checkTriNum(nums, triNums)==1){
                        flag = 1;
                        break;
                    }
                }
                if(flag == 1){
                    break;
                }
            }
            System.out.println(flag);
        }
    }

    private static int[] calculateNum(int k, int p, int q) {
        int[] ret = new int[3];
        int[] arr = new int[k];
        for(int i=0;i<k;i++){
            arr[i] = 1;
        }
        for(int i=0;i<p+1;i++){
            ret[0] += 1;
        }
        for(int i=p+1;i<q+1;i++){
            ret[1] += 1;
        }
        for(int i=q+1;i<k;i++){
            ret[2] += 1;
        }
        return ret;
    }

    private static int checkTriNum(int[] nums, HashSet triArr) {
        int flag = 1;
        for(int i=0;i<3;i++){
            int x = nums[i];
            if(!triArr.contains(x)){
                flag = 0;
                return flag;
            }
        }
        return flag;
    }
}
