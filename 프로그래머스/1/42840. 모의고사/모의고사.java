import java.util.ArrayList;
class Solution {
    public int[] solution(int[] answers) {
        
        int[] score = {0, 0, 0};
        int[] first = {1, 2, 3, 4, 5};
        int[] second = {2, 1, 2, 3, 2, 4, 2, 5};
        int[] third = {3, 3, 1, 1, 2, 2, 4, 4, 5, 5};
        
        int maxValue = 0;
        
        for (int i=0; i<answers.length; i++){
            if (answers[i]==first[i%first.length]){
                score[0]++;

            }
            if (answers[i]==second[i%second.length]){
                score[1]++;
            }
            if (answers[i]==third[i%third.length]){
                score[2]++;
            }
            for(int j=0;j<3;j++){
                if(maxValue < score[j]){
                    maxValue = score[j];
                }
            }
        }
        ArrayList<Integer> arr = new ArrayList<Integer>();
        for(int i=0;i<3;i++){
            if(maxValue==score[i]){
                arr.add(i+1);
            }
        }
        int[] answer = arr.stream()
                .mapToInt(Integer::intValue)
                .toArray();
        return answer;
    }
}