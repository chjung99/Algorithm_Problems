class Solution {
    public boolean wordPattern(String pattern, String s) {
        String[] tokens = s.split(" ");
        Map<Character, String> map = new HashMap<>();
        Set<String> used = new HashSet<>();

        if (pattern.length() != tokens.length) {
            return false;
        }

        for (int i = 0; i < pattern.length(); i++) {
            Character key = pattern.charAt(i);

            if (map.containsKey(key) && !map.get(key).equals(tokens[i])) {
                return false;
            }
            if (!map.containsKey(key)) {
                if (used.contains(tokens[i])){
                    return false;
                }
                map.put(key, tokens[i]);
                used.add(tokens[i]);
            }
        }
        return true;
    }
}