import java.util.TreeSet;

class Solution {
    public int[] solution(int[] numbers) {
        
        TreeSet<Integer> set = new TreeSet<>();
        for(int i=0;i<numbers.length;i++){
            for(int j=0;j<numbers.length;j++){
                if(i==j)continue;
                set.add(numbers[i]+numbers[j]);
            }
        }
        int[] answer = new int[set.size()];
        for(int i=0;i<answer.length;i++){
            answer[i] = set.pollFirst();
        }
        
        return answer;
    }
}