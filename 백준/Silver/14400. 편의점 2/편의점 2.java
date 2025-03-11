import java.io.*;
import java.util.*;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));
        StringTokenizer st = new StringTokenizer(br.readLine());

        int N = Integer.parseInt(st.nextToken());

        int[] x = new int[N];
        int[] y = new int[N];

        for(int i=0; i<N; i++){
            st = new StringTokenizer(br.readLine());
            x[i] = Integer.parseInt(st.nextToken());
            y[i] = Integer.parseInt(st.nextToken());
        }

        Arrays.sort(x);
        Arrays.sort(y);

        //편의점 위치
        int midX = x[N/2];
        int midY = y[N/2];

        long distX = 0;
        long distY = 0;

        for(int i=0; i<N; i++){
            distX += Math.abs(midX - x[i]);
            distY += Math.abs(midY - y[i]);
        }

        long sum = distX + distY;
        System.out.print(sum);
    }
}