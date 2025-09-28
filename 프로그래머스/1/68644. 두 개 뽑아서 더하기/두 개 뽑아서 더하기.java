import java.util.*;

class Solution {
    public int[] solution(int[] numbers) {
        Set<Integer> set = new HashSet<Integer>();
        for (int i=0; i<numbers.length; i++){
            for (int j=i+1; j<numbers.length; j++){
                set.add(numbers[i]+numbers[j]);
            }
        }
        int[] answer = set.stream()
            .mapToInt(x -> x.intValue())
            .toArray();
        
        Arrays.sort(answer);
        return answer;
    }
}