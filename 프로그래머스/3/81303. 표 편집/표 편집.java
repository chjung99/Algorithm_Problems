import java.util.*;


class Solution {
    public String solution(int n, int k, String[] cmd) {
        int cur = k;
        
        int[] up = new int[n];
        int[] down = new int[n];
        
        Deque<List<Integer>> deque = new ArrayDeque<>();
        
        for (int i=0; i<n ;i++){
            up[i] = i-1;
            down[i] = i+1;
        }
        
         for (String s: cmd){
            String[ ] sp = s.split(" ");
             
            if (s.length()>1){
                // int d = Character.getNumericValue(s.charAt(2));
                
                int d = Integer.parseInt(sp[1]);
                if (sp[0].equals("D")){
                    for (int i=0; i<d ;i++){
                        cur = down[cur];
                    }
                } else{
                    for (int i=0; i<d ;i++){
                        cur = up[cur];
                    }
                }


            } else{

                if (sp[0].equals("C")){
                    deque.push(Arrays.asList(cur, up[cur], down[cur]));
                    if (up[cur]==-1){//최상단인 경우

                        up[down[cur]] = up[cur];
                        cur = down[cur];
                        
                    }else if(down[cur]==n){//최하단인 경우
                        
                        down[up[cur]] = down[cur];
                        cur = up[cur];
                        
                    }else{
                        
                        down[up[cur]] = down[cur];
                        up[down[cur]] = up[cur];
                        cur = down[cur];
                        
                    }

                } else{ // 'Z'
                    List<Integer> prevList = deque.pop();

                    int prev = prevList.get(0);
                    int prevUp = prevList.get(1);
                    int prevDown = prevList.get(2);

                    if (prevUp == -1){//최상단인 경우
                        up[prevDown] = prev;
                    } else if(prevDown == n){//최하단인 경우
                        down[prevUp] = prev;
                    } else{
                        up[prevDown] = prev;
                        down[prevUp] = prev;
                    }


                }
            }
        }
        String answer = "O".repeat(n);
        StringBuilder sb = new StringBuilder(answer);
        while (!deque.isEmpty()){
            List<Integer> list = deque.pop();
            sb.setCharAt(list.get(0), 'X');
        }
        answer = sb.toString();
        return answer;
    }
}