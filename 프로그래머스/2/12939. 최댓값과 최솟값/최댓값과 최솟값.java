import java.util.*;

class Solution {
    public String solution(String s) {
        String answer = "";
        String[] strArray = s.split(" ");
        int max = Arrays.stream(strArray)
                .mapToInt(Integer::parseInt)
                .max().orElseThrow();

        int min = Arrays.stream(strArray)
                .mapToInt(Integer::parseInt)
                .min().orElseThrow();
        
        return min + " " + max;
    }
}