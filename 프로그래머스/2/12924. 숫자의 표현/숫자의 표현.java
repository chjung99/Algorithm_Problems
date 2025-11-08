class Solution {
    public int solution(int n) {
        
        int answer = 0;
        
        for (int i = 1; i <= n; i++){
            if (findContinuousNumber(i, n) == n) answer++;
        }
        
        return answer;
    }
    
    public int findContinuousNumber(int start, int n) {
        int sum = 0;
        for (int i = start ; i <= n; i++) {
            if (sum >= n){
                break;
            }
            sum += i;
        }
        return sum;
    }
}