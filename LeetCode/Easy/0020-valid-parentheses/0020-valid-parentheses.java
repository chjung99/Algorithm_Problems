class Solution {
    public boolean isValid(String s) {
        Deque<Character> st = new ArrayDeque<>();
        int n = s.length();

        for (int i = 0; i < n; i++) {
            Character key = s.charAt(i);
            if (key == '(' || key == '{' || key == '[') {
                st.add(key);
                continue;
            }

            if (st.isEmpty()) {
                return false;
            }

            if (key == ')' && st.peekLast() == '(') {
                st.removeLast();
                continue;
            }

            if (key == '}' && st.peekLast() == '{') {
                st.removeLast();
                continue;
            }

            if (key == ']' && st.peekLast() == '[') {
                st.removeLast();
                continue;
            }
            return false;
        }
        return st.size() == 0;

    }
}