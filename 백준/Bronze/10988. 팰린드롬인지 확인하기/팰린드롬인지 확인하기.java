import java.util.Scanner;

public class Main
{
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String s = sc.nextLine();
        StringBuilder x = new StringBuilder();
        for (int i=s.length()-1;i >= 0 ;i--){
            x.append(s.charAt(i));
        }
        if (s.contentEquals(x)){
            System.out.println("1");
        }
        else {
            System.out.println("0");
        }
    }
}
