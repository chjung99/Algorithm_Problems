class Solution {
    public int countSubstrings(String s) {
        int answer = 0;
        int n = s.length();

        for (int i = 0; i < n; i++) {
            for (int j = i+1; j <= n; j++) {
                answer += isPalindrome(s.substring(i, j));
            }
        }
        return answer;
    }

    public int isPalindrome(String s) {
        StringBuilder sb = new StringBuilder(s);
        if (sb.reverse().toString().equals(s)) return 1;
        return 0;
    }
}