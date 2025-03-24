import java.util.Scanner;

public class Main
{
	public static void main(String[] args) {
	    Scanner sc = new Scanner(System.in);
	    
		while (sc.hasNextLine()){
		    String a = sc.nextLine();
		    String b = sc.nextLine();
		    
		    int[] aArr = new int[26];
		    int[] bArr = new int[26];
		    
		    for (int i = 0; i < a.length(); i++){
		        int idx = a.charAt(i)-'a';
		        aArr[idx] += 1;
		    }
		    
		    for (int i = 0; i < b.length(); i++){
		        int idx = b.charAt(i)-'a';
		        bArr[idx] += 1;
		    }
		    
		    for (int i = 0 ; i<26;i++){
		        if(aArr[i]>=1 && bArr[i]>=1){
		            int cnt = Math.min(aArr[i],bArr[i]);
		            for(int j=0;j<cnt;j++){
		                System.out.print((char)('a'+i));
		            }
		        }
		    }
		    System.out.println("");
		    
		}
		
	}
}

