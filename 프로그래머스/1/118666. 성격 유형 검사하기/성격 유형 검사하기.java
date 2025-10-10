import java.util.*;

class Solution {
    public String solution(String[] survey, int[] choices) {
        StringBuilder sb = new StringBuilder();
        String answer = "";
        int n = survey.length;
        HashMap<String, Integer> hashMap = new HashMap<>();
        String[][] arr = {{"R", "T"}, {"C","F"}, {"J","M"}, {"A","N"}};
        for (int i=0; i<4; i++){
            hashMap.put(arr[i][0], 0);
            hashMap.put(arr[i][1], 0);
        }
        
        for (int i=0; i<n; i++){
            String disagree = String.valueOf(survey[i].charAt(0));
            String agree = String.valueOf(survey[i].charAt(1));
            int score = choices[i] - 4;
            
            if (score > 0){ // 동의
                hashMap.put(agree, hashMap.get(agree)+score);
            } else if (score < 0){
                
                hashMap.put(disagree, hashMap.get(disagree)+Math.abs(score));
            }
        }
        
        for (int i=0; i<4; i++){
            String firstKey = arr[i][0];
            String secondKey = arr[i][1];
            if (hashMap.get(firstKey)>hashMap.get(secondKey)){
                sb.append(firstKey);
            } else if (hashMap.get(firstKey)<hashMap.get(secondKey)){
                sb.append(secondKey);
            } else {
                if (firstKey.compareTo(secondKey)<0){
                    sb.append(firstKey);
                } else {
                    sb.append(secondKey);
                }
            }
        }
        
        return sb.toString();
    }
}