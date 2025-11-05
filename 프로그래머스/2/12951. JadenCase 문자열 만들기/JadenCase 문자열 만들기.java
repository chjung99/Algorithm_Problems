import java.util.*;

class Solution {
    public String solution(String s) {
        StringBuilder sb = new StringBuilder();
        String[] split = s.split(" ");
        
        for (int i=0; i<split.length; i++){
            String str = split[i];
            
            if (str.length()==0){
                sb.append(" ");
                continue;
            }
            
            if (isAlphabet(str.charAt(0))){
                sb.append(str.substring(0, 1).toUpperCase());
                if (str.length() > 1) {
                    sb.append(str.substring(1).toLowerCase());
                }
            } else {
                sb.append(str.substring(0, 1));
                if (str.length() > 1) {
                    sb.append(str.substring(1).toLowerCase());
                }
            }
            if (i != split.length-1){
                sb.append(" ");
            }
        }
        if (sb.toString().length() < s.length()){
            sb.append(" ");
        }
        
        System.out.print("\""+sb.toString()+"\"");
        
        return sb.toString();
    }
    public boolean isAlphabet(char c) {
        return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z');
    }
}