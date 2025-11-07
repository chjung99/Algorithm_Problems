import java.util.*;

class Solution {
    public int[] solution(String s) {
        int[] answer = {};
        int binaryTransformCount = 0;
        int reducedZeroNum = 0;
        String reducedZero;
        String transformed;
        
        while (true) {
            if (s.equals("1")) break;
            reducedZero = reduceZero(s);
            reducedZeroNum += s.length() - reducedZero.length();
            
            s = toBinary(reducedZero.length());
            binaryTransformCount += 1;
        }
        
        return new int[]{binaryTransformCount, reducedZeroNum};
    }
    public String toBinary(int c) {
        StringBuilder sb = new StringBuilder();
        
        while (c != 0){
            sb.append(Integer.toString(c%2));
            c = (int) c / 2;
        }
        return sb.reverse().toString();
    }
    public String reduceZero(String s){
        StringBuilder sb = new StringBuilder();
        
        for (int i=0; i < s.length(); i++){
            if (s.charAt(i)=='0') continue;
            sb.append(Character.toString(s.charAt(i)));
        }
        return sb.toString();
    }
}