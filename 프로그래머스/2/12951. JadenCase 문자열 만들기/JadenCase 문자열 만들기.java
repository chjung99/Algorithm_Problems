import java.util.*;

class Solution {
    public String solution(String s) {
        StringBuilder sb = new StringBuilder();
        String[] split = s.toLowerCase().split("");
        boolean flag = isAlphabet(s.charAt(0));
        for (String str: split){
            String token = flag ? str.toUpperCase() : str;
            flag = str.equals(" ");
            sb.append(token);
        }
        
        
        return sb.toString();
    }
    public boolean isAlphabet(char c) {
        return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z');
    }
}