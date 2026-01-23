class Solution {
    public int romanToInt(String s) {
        int n = s.length();
        Deque<Character> deque = new ArrayDeque<>();
        Map<Character, Integer> map = new HashMap<>();
        map.put('I', 1);
        map.put('V', 5);
        map.put('X', 10);
        map.put('L', 50);
        map.put('C', 100);
        map.put('D', 500);
        map.put('M', 1000);

        int answer = 0;
        for (int i = 0; i < n - 1; i++) {
            Character key = s.charAt(i);
            Character nextKey = s.charAt(i+1);
            if (map.get(key) >= map.get(nextKey)) {
                answer += map.get(key);
            } else {
                answer -= map.get(key);
            }
        }
        answer += map.get(s.charAt(n-1));
        return answer;
    }
}