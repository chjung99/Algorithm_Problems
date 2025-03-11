import java.util.Scanner;
import java.util.HashSet;
import java.util.StringTokenizer;
public class Main
{
	public static void main(String[] args) {
	    Scanner sc = new Scanner(System.in);
		int cnt = 0;
		String str = sc.nextLine();
		StringTokenizer st = new StringTokenizer(str);
		while (st.hasMoreTokens()){
		    String key = st.nextToken();
		    cnt += 1;
		    
		}
		System.out.print(cnt);
	}
}
