/******************************************************************************

 Welcome to GDB Online.
 GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
 C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
 Code, Compile, Run and Debug online from anywhere in world.

 *******************************************************************************/
import java.util.*;
import java.util.Collections;

public class Main
{
    private static class Circle{
        int type;
        int value;
        int idx;
        Circle(int type, int value, int idx){
            this.type = type;
            this.value = value;
            this.idx = idx;
        }
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int N = sc.nextInt();
        Stack<Circle> st = new Stack<Circle>();
        ArrayList<Circle> arr = new ArrayList<Circle>();
        boolean status = true;
        for(int i=0; i < N; i++){
            int x = sc.nextInt();
            int r = sc.nextInt();
            arr.add(new Circle(0,x-r, i));
            arr.add(new Circle(1, x+r, i));
        }
        Collections.sort(arr, new Comparator<Circle>() {
            @Override
            public int compare(Circle o1, Circle o2) {
                return Integer.compare(o1.value, o2.value);
            }
        });

        for (int i=0; i < arr.size(); i++){
            Circle cur = arr.get(i);
            Circle nxt = null;
            if (i+1 < arr.size()) {
                nxt = arr.get(i + 1);
            }

            if (st.isEmpty()){
                st.push(cur);
            }
            else {
                Circle top = st.peek();
                if (cur.type == 0){
                    if (top.value < cur.value){
                        st.push(cur);
                    }
                    else{
                        status = false;
                        break;
                    }
                }
                if (cur.type == 1){
                    if (top.idx == cur.idx){
                        if(nxt != null && cur.value == nxt.value){
                            status = false;
                            break;
                        }
                        st.pop();
                    }
                    else{
                        status = false;
                        break;
                    }
                }

            }
        }
        if (status){
            System.out.print("YES");
        }
        else{
            System.out.print("NO");
        }
    }
}
